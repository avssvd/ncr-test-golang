FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["${dbuser}", "${dbpass}", "${dbname}", "${dburi}", "${dbport}", "${grpcport}", "${restport}"]
ENTRYPOINT ["./app"]
