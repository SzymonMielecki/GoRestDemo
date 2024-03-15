BINARY_NAME=air_qual
export PORT=:1323
include .env

build_client:
	@go build -o ${BINARY_NAME}_client client/main.go

build_server:
	@go build -o ${BINARY_NAME}_server server/main.go

build: setup build_client build_server

run_server:
	./air_qual_server

run_client:
	./air_qual_client -k ${API_KEY}

run: run_server run_client

full_server: build_server run_server

full_client: build_client run_client

target: full_server full_client

test:
	@go test -v ./...

