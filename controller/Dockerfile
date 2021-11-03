FROM golang:1.17 AS service-dependencies
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM service-dependencies AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.14
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/app ./app
ENTRYPOINT ./app --serial=${serial} --servuri=${servuri} --servport=${servport}
