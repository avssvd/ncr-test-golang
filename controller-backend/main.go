package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/avssvd/ncr-test-golang/gen/rest/restapi"
	"github.com/avssvd/ncr-test-golang/gen/rest/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"strconv"
	"time"

	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	report "github.com/avssvd/ncr-test-golang/gen/proto/go/api/grpc/report"
	db "github.com/avssvd/ncr-test-golang/gen/sqlc"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type Options struct {
	DBUser   string
	DBPass   string
	DBName   string
	DBURI    string
	DBPort   int
	GRPCPort int
	RESTPort int
}

func (opts *Options) get() {
	flag.StringVar(&opts.DBUser, "dbuser", "app", "database user")
	flag.StringVar(&opts.DBPass, "dbpass", "pass", "database password")
	flag.StringVar(&opts.DBName, "dbname", "app", "database name")
	flag.StringVar(&opts.DBURI, "dburi", "db", "database URI")
	flag.IntVar(&opts.DBPort, "dbport", 5432, "database port")
	flag.IntVar(&opts.GRPCPort, "grpcport", 8080, "port to listen on gRPC")
	flag.IntVar(&opts.RESTPort, "restport", 8081, "port to listen on REST")

	flag.Parse()
}

func (opts *Options) check() error {
	const (
		minPort = 1
		maxPort = 65535
	)

	errList := make([]string, 0, 5)

	if opts.DBPort < minPort || opts.DBPort > maxPort {
		errList = append(errList, fmt.Sprintf("database port %d out of range [%d:%d]\n", opts.DBPort, minPort, maxPort))
	}
	if opts.GRPCPort < minPort || opts.GRPCPort > maxPort {
		errList = append(errList, fmt.Sprintf("gRPC port %d out of range [%d:%d]\n", opts.GRPCPort, minPort, maxPort))
	}
	if opts.RESTPort < minPort || opts.RESTPort > maxPort {
		errList = append(errList, fmt.Sprintf("REST port %d out of range [%d:%d]\n", opts.RESTPort, minPort, maxPort))
	}
	if opts.GRPCPort == opts.RESTPort {
		errList = append(errList, fmt.Sprintf("REST and gRPC ports are the same: %d\n", opts.GRPCPort))
	}
	if len(errList) != 0 {
		return fmt.Errorf(strings.Join(errList, ""))
	}

	return nil
}

func main() {
	var opts Options

	opts.get()

	if err := opts.check(); err != nil {
		log.Fatal(err)
	}

	connstr := fmt.Sprintf("dbname=%s host=%s port=%d user=%s password='%s' sslmode=disable",
		opts.DBName, opts.DBURI, opts.DBPort, opts.DBUser, opts.DBPass)
	database, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal("failed to open db:", err.Error())
	}

	queries := db.New(database)


	go grpcServe(opts.GRPCPort, queries)

	restServe(opts.RESTPort, queries)
}

func grpcServe(grpcPort int, db *db.Queries){
	listenOn := fmt.Sprintf(":%d", grpcPort)
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		log.Fatalf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	report.RegisterReportServiceServer(server, &reportServiceServer{db: db})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		log.Fatal("failed to serve gRPC server: %w", err)
	}
}

type reportServiceServer struct {
	report.UnimplementedReportServiceServer
	db *db.Queries
}

func (s *reportServiceServer) PutReport(ctx context.Context, req *report.PutReportRequest) (*report.PutReportResponse, error) {
	serial := req.GetSerial()
	indication := req.GetIndication()

	err := s.db.CreateIndication(ctx, db.CreateIndicationParams{
		Indication:       fmt.Sprintf("%.1f", indication),
		ControllerSerial: serial,
		SentAt:           time.Now(),
	})

	if err != nil {
		return &report.PutReportResponse{
			// TODO generate? timeBeforeNextConnInSec
			TimeBeforeNextConnInSec: 10,
			ErrorMessage:            err.Error(),
		}, nil

	}

	log.Printf("New report from %v: temperature is %v\n", serial, indication)

	return &report.PutReportResponse{
		TimeBeforeNextConnInSec: 10,
		ErrorMessage:            "",
	}, nil
}

func restServe(serverPort int, db *db.Queries) {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewBackendAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = serverPort

	api.GetControllersHandler = operations.GetControllersHandlerFunc(
		func(params operations.GetControllersParams) middleware.Responder {
			controllers, err := db.ListControllers(context.Background())
			if err != nil {
				log.Println(err)
				return operations.NewGetControllersOK()
			}
			
			body := operations.GetControllersOKBody{}
			for _, controller := range controllers {
				body.Controllers = append(body.Controllers, &operations.GetControllersOKBodyControllersItems0{
					CreatedAt: strfmt.DateTime(controller.CreatedAt),
					Serial:    controller.Serial,
				})
			}

			return operations.NewGetControllersOK().WithPayload(&body)
		})

	api.GetControllerIndicationsHandler = operations.GetControllerIndicationsHandlerFunc(
		func(params operations.GetControllerIndicationsParams) middleware.Responder {
			_, err := db.GetController(context.Background(), params.Controller.Serial)
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return operations.NewGetControllerIndicationsBadRequest().WithPayload(&operations.GetControllerIndicationsBadRequestBody{Error: "controller not found"})
			case err != nil:
				log.Println(err)
				return operations.NewGetControllerIndicationsBadRequest().WithPayload(&operations.GetControllerIndicationsBadRequestBody{Error: err.Error()})
			}

			indications, err := db.ListIndicationsByController(context.Background(), params.Controller.Serial)
			if err != nil {
				log.Println(err)
				return operations.NewGetControllerIndicationsBadRequest().WithPayload(&operations.GetControllerIndicationsBadRequestBody{Error: err.Error()})
			}

			body := operations.GetControllerIndicationsOKBody{}
			for _, indication := range indications {
				temp, err := strconv.ParseFloat(indication.Indication, 32)
				if err != nil {
					log.Println(err)
				}
				body.Indications = append(body.Indications, &operations.GetControllerIndicationsOKBodyIndicationsItems0{
					Indication: float32(temp),
					SentAt:     strfmt.DateTime(indication.SentAt),
				})
			}

			return operations.NewGetControllerIndicationsOK().WithPayload(&body)
		})

	api.PostControllerHandler = operations.PostControllerHandlerFunc(
		func(params operations.PostControllerParams) middleware.Responder {
			_, err := db.GetController(context.Background(), params.Controller.Serial)
			switch {
			case err == nil:
				return operations.NewPostControllerBadRequest().WithPayload(&operations.PostControllerBadRequestBody{Error: "controller already exists"})
			case !errors.Is(err, sql.ErrNoRows):
				log.Println(err)
				return operations.NewPostControllerBadRequest().WithPayload(&operations.PostControllerBadRequestBody{Error: err.Error()})
			}

			err = db.CreateController(context.Background(), params.Controller.Serial)
			if err != nil {
				log.Println(err)
				return operations.NewPostControllerBadRequest().WithPayload(&operations.PostControllerBadRequestBody{Error: err.Error()})
			}

			return operations.NewPostControllerOK().WithPayload(&operations.PostControllerOKBody{Success: true})
		},
	)

	api.DeleteControllerHandler = operations.DeleteControllerHandlerFunc(
		func(params operations.DeleteControllerParams) middleware.Responder {
			affectedRows, err := db.DeleteController(context.Background(), params.Controller.Serial)
			if err != nil {
				log.Println(err)
				return operations.NewDeleteControllerBadRequest().WithPayload(&operations.DeleteControllerBadRequestBody{Error: err.Error()})
			}
			if affectedRows == 0 {
				return operations.NewDeleteControllerBadRequest().WithPayload(&operations.DeleteControllerBadRequestBody{Error: "controller not found"})
			}
			return operations.NewDeleteControllerOK().WithPayload(&operations.DeleteControllerOKBody{Success: true})
		},
	)


	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}