# Controller

---
* [Controller-backend](#controller-backend)
* [Controller](#controller)

---

## Controller-backend

[REST API docs](https://app.swaggerhub.com/apis/avssvd/ncr-test-golang-openapi-2.0/1.0.0)

[swagger.yaml](api/rest/swagger.yaml)

[report.proto](api/grpc/report/report.proto)

[DB schema](db/sqlc/migration/000001_init_schema.up.sql)

### Run 

```shell
cd controller-backend
go mod download
go run main.go -dbuser=root -dbpass=pass -dbname=app -dburi=127.0.0.1 -dbport=5432 -grpcport=8000 -restport=8001
```

### Docker

Создать и отредактировать файл .env на основе .env.example:

```shell
cp .env.example .env
```

Собрать образ ncr-controller-backend и скачать образ PostgreSQL:

```shell
make
```

Создать и запустить контейнеры:

```shell
make run-server
```

Произвести миграцию
```shell
make migrate
```

## Controller

### Run

```shell
cd controller
go mod download
go run controller/main.go --serial=abcdefg12345 --servuri=127.0.0.1 --servport=8000
```

### Docker
Собрать образ ncr-controller:

```shell
make
```

Создать и запустить контейнеры:

```shell
make run-client ARGS="--serial=abcdefg12345 --servuri=server --servport=8000"
```