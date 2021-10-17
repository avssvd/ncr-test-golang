package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	reportv1 "github.com/avssvd/ncr-test-golang/gen/proto/go/api/grpc/report"
	"google.golang.org/grpc"
)

const (
	minTemp       = 20
	maxTemp       = 40
	errDelayInSec = 6
	maxNFail      = 3
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

func (opts *Options) get() {
	flag.StringVar(&opts.Serial, "serial", "unknown", "controller serial")
	flag.StringVar(&opts.ServerURI, "servuri", "127.0.0.1", "server URI")
	flag.IntVar(&opts.ServerPort, "servport", 8080, "server port")

	flag.Parse()
}

// win 	-> 	message + wait TimeBeforeNextConnInSec
// err 	-> 	message + wait 1 min for next conn
//			3x err in a row -> shutdown

func main() {
	var opts Options

	opts.get()

	if err := opts.check(); err != nil {
		log.Fatal(err)
	}

	nFail := 0
	for nFail < maxNFail {
		delayInSec, err := sendReport(opts.ServerURI, opts.ServerPort, opts.Serial)
		if err != nil {
			log.Println(err)
			nFail++
			if nFail < maxNFail {
				time.Sleep(errDelayInSec * time.Second)
			}
			continue
		}
		nFail = 0
		time.Sleep(time.Duration(delayInSec) * time.Second)
	}
}

func sendReport(servURI string, servPort int, serial string) (delayInSec int64, err error) {
	connectTo := fmt.Sprintf("%s:%d", servURI, servPort)
	conn, err := grpc.Dial(connectTo, grpc.WithInsecure())
	if err != nil {
		return 0, fmt.Errorf("failed to connect to ReportService on %s: %w", connectTo, err)
	}

	controller := reportv1.NewReportServiceClient(conn)
	resp, err := controller.PutReport(context.Background(), &reportv1.PutReportRequest{
		Serial:     serial,
		Indication: getTemp(),
	})

	switch {
	case err != nil:
		return 0, fmt.Errorf("failed to PutReport: %w", err)

	case resp.ErrorMessage != "":
		return 0, fmt.Errorf("server error: %s", resp.ErrorMessage)
	}

	log.Printf("Successfully PutReport; %d sec for next report", resp.TimeBeforeNextConnInSec)
	return resp.TimeBeforeNextConnInSec, nil
}

func getTemp() float32 {
	rand.Seed(time.Now().Unix())
	return minTemp + rand.Float32()*(maxTemp-minTemp)
}
