package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	reportv1 "github.com/avssvd/ncr-test-golang/gen/proto/go/reportapis/report/v1"
	"google.golang.org/grpc"
)


type Options struct {
	DBUser   string
	DBPass   string
	DBURI    string
	DBPort   int
	GRPCPort int
	HTTPPort int
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
	if opts.HTTPPort < minPort || opts.HTTPPort > maxPort {
		errList = append(errList, fmt.Sprintf("HTTP port %d out of range [%d:%d]\n", opts.HTTPPort, minPort, maxPort))
	}
	if opts.GRPCPort == opts.HTTPPort {
		errList = append(errList, fmt.Sprintf("HTTP and gRPC ports are the same: %d\n", opts.GRPCPort))
	}
	if len(errList) != 0 {
		return fmt.Errorf(strings.Join(errList, ""))
	}

	return nil
}

func main() {
	var opts Options
	flag.StringVar(&opts.DBUser, "dbuser", "app", "database user")
	flag.StringVar(&opts.DBPass, "dbpass", "pass", "database password")
	flag.StringVar(&opts.DBURI, "dburi", "db", "database URI")
	flag.IntVar(&opts.DBPort, "dbport", 5432, "database port")
	flag.IntVar(&opts.GRPCPort, "grpcport", 8080, "port to listen on gRPC")
	flag.IntVar(&opts.HTTPPort, "httpport", 8081, "port to listen on HTTP")

	flag.Parse()

	if err := opts.check(); err != nil {
		log.Fatal(err)
	}

	if err := grpcServe(opts.GRPCPort); err != nil {
		log.Fatal(err)
	}
}

func grpcServe(grpcPort int) error {
	listenOn := fmt.Sprintf(":%d", grpcPort)
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	reportv1.RegisterReportServiceServer(server, &reportServiceServer{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type reportServiceServer struct {
	reportv1.UnimplementedReportServiceServer
}

func (s *reportServiceServer) PutReport(ctx context.Context, req *reportv1.PutReportRequest) (*reportv1.PutReportResponse, error) {
	serial := req.GetSerial()
	indication := req.GetIndication()

	log.Printf("New report from %v: temperature is %v\n", serial, indication)

	return &reportv1.PutReportResponse{
		TimeBeforeNextConnInSec: 10,
		ErrorMessage:            "",
	}, nil
}
