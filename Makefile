all: pull build

pull:
	docker pull postgres:14.0
	docker pull adminer:4
	docker pull migrate/migrate:4

build: build-server build-client

build-client:
	docker build -f ./controller/Dockerfile -t ncr-controller:latest ./controller/

build-server:
	docker build -f ./controller-backend/Dockerfile -t ncr-controller-backend:latest ./controller-backend/

run-client:
	docker run --rm --network $$(basename $$(readlink -e $$(pwd)))_ncr-network ${ARGS} ncr-controller:latest

run-server:
	docker-compose up -d

stop-server:
	docker-compose stop

migrate:
	docker run --rm -v $$(readlink -e ./db/sqlc/migration):/migrations --network $$(basename $$(readlink -e $$(pwd)))_ncr-network migrate/migrate -path=/migrations/ -database postgres://root:pass@db:5432/app?sslmode=disable up

clean:
	for client in $$(docker ps -a | grep "ncr-controller:latest" | awk '{print $$1}') ; do \
		docker stop $$client ; \
	done
	docker-compose down
	sudo rm -rf _postgres-data
	docker image rm ncr-controller:latest
	docker image rm ncr-controller-backend:latest
	docker image rm postgres:14.0
	docker image rm adminer:4
	docker image rm migrate/migrate:4