package main

import (
	"context"
	"flag"
	"fmt"
	reportv1 "github.com/avssvd/ncr-test-golang/gen/proto/go/reportapis/report/v1"
	"google.golang.org/grpc"
	"log"
)

type Options struct {
	Serial     string
	ServerURI  string
	ServerPort int
}

func (opts *Options) check() error {
	const (
		minPort = 1
		maxPort = 65535
	)

	if opts.ServerPort < minPort || opts.ServerPort > maxPort {
		return fmt.Errorf("server port %d out of range [%d:%d]\n", opts.ServerPort, minPort, maxPort)
	}

	return nil
}

func main() {
	var opts Options
	flag.StringVar(&opts.Serial, "serial", "unknown", "controller serial")
	flag.StringVar(&opts.ServerURI, "servuri", "127.0.0.1", "server URI")
	flag.IntVar(&opts.ServerPort, "servport", 8080, "server port")

	flag.Parse()

	if err := opts.check(); err != nil {
		log.Fatal(err)
	}
	if err := sendReport(opts.ServerURI, opts.ServerPort, opts.Serial); err != nil {
		log.Fatal(err)
	}
}
func sendReport(servURI string, servPort int, serial string) error {
	connectTo := fmt.Sprintf("%s:%d", servURI, servPort)
	fmt.Println(connectTo)
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to ReportService on %s: %w", connectTo, err)
	}
	log.Println("Connected to", connectTo)

	controller := reportv1.NewReportServiceClient(conn)
	if _, err := controller.PutReport(context.Background(), &reportv1.PutReportRequest{
		Serial:     serial,
		Indication: 40.4,
	}); err != nil {
		return fmt.Errorf("failed to PutReport: %w", err)
	}

	log.Println("Successfully PutReport")
	return nil
}
