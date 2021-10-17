package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"

	db "github.com/avssvd/ncr-test-golang/gen/sqlc"
	_ "github.com/lib/pq"
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
