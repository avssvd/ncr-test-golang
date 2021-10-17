all: prepare

prepare: build
	docker pull postgres:14.0

build: build-server build-client

build-client:
	docker build -f ./controller/Dockerfile -t ncr-controller ./controller/

build-server:
	docker build -f ./controller-backend/Dockerfile -t ncr-controller-backend ./controller-backend/

run-client:
	docker run --rm --network ncr_ncr-network ncr-controller:latest ${ARGS}

run-server:
	docker-compose run ${ARGS}

migrate:
	migrate -path ./db/sqlc/migration -database postgres://root:pass@127.0.0.1:5432/app?sslmode=disable up