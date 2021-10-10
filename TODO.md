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
- [ ] create report.proto
- [ ] create buf.yaml
- [ ] create buf.gen.yaml

### http
- [ ] create swagger.yaml 

## Controller-backend
- [ ] create gRPC server
- [ ] create http server
- [ ] create functions 4 working with db

## Controller
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
