.DEFAULT_GOAL := build

.PHONY:vet build

vet:
		go vet ./...
proto:
	protoc ./pkg/**/pb/*.proto --go_out=. --go-grpc_out=.

server:
	go run ./cmd/

build: vet
		go build -o ./bin/api ./cmd/
