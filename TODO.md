# TODO:

## Docker (Docker-compose)
- [ ] add Dockerfile 4 controller-backend
- [ ] add Dockerfile 4 controller
- [ ] add docker-compose 4 deploy controller-backend & PostgreSQL

## DB
- [ ] create schema
- [ ] create migrations
- [ ] create methods for User
- [ ] create methods for IPCheck

## API
### gRPC
- [x] create report.proto
- [x] create buf.yaml
- [x] create buf.gen.yaml

### http
- [ ] create swagger.yaml 

## Controller-backend
- [x] add command-line flag parsing
- [ ] create gRPC server
- [ ] create http server
- [ ] create functions 4 working with db

## Controller
- [ ] add command-line flag parsing
- [ ] create gRPC client
- [ ] create scheduler 4 report sending

## Makefile
- [ ] add build (build-server + build-client)
- [ ] add build-server
- [ ] add build-client
- [ ] add run (run-server + run-client)
- [ ] add run-server
- [ ] add run-client
- [ ] add stop (stop-server + stop-client)
- [ ] add stop-server
- [ ] add stop-client
- [ ] add clean (stop + docker image rm ...)
