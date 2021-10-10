package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	reportv1 "github.com/avssvd/ncr-test-golang/gen/proto/go/reportapis/report/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
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

type reportServiceServer struct{
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
