FROM golang:1.17
WORKDIR /go/src/github.com/avssvd/ncr-test-golang/controller
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/avssvd/ncr-test-golang/controller/app ./app
ENTRYPOINT ["./app"]
