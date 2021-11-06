# TODO:

## Fixes
- [x] hardcode? internal db port
- [x] add <code>adminer</code> service to network
- [x] use environment variables instead of dotenv file
- [x] remove extra <code>.go</code> files from backend image
- [x] use <code>docker-compose build</code>
- [x] make migrations via <code>docker-entrypoint-initdb.d</code>
- [x] use <code>db.Ping()</code> 4 check db connection
- [x] add logging http requests
- [ ] rewrite db client on GORM

## Docker (Docker-compose)
- [x] add Dockerfile 4 controller-backend
- [x] add Dockerfile 4 controller
- [x] add docker-compose 4 deploy controller-backend & PostgreSQL

## DB
- [x] create schema
- [x] create migrations
- [x] create methods for User
- [x] create methods for IPCheck

## API
### gRPC
- [x] create report.proto
- [x] create buf.yaml
- [x] create buf.gen.yaml

### http
- [x] create swagger.yaml 
- [x] generate server via go-swagger

## Controller-backend
- [x] add command-line flag parsing
- [x] create gRPC server
- [x] create http server
- [x] create functions 4 working with db

## Controller
- [x] add command-line flag parsing
- [x] create gRPC client
- [x] create scheduler 4 report sending

## Makefile
- [x] add build (build-server + build-client)
- [x] add build-server
- [x] add build-client
- [x] add run-server
- [x] add run-client
- [x] add clean (stop + docker image rm ...)
