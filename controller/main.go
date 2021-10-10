package main

import (
	"context"
	"fmt"
	"log"

	reportv1 "github.com/avssvd/ncr-test-golang/gen/proto/go/reportapis/report/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to ReportService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	controller := reportv1.NewReportServiceClient(conn)
	if _, err := controller.PutReport(context.Background(), &reportv1.PutReportRequest{
		Serial:     "avssvd",
		Indication: 40.4,
	});	err != nil {
		return fmt.Errorf("failed to PutReport: %w", err)
	}

	log.Println("Successfully PutReport")
	return nil
}
