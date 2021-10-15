package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	report "github.com/avssvd/ncr-test-golang/gen/proto/go/api/grpc/report"
	db "github.com/avssvd/ncr-test-golang/gen/sqlc"
	"google.golang.org/grpc"
)

func grpcServe(grpcPort int, db *db.Queries) {
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
