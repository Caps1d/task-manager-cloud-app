.DEFAULT_GOAL := build

.PHONY:vet build

vet:
		go vet ./...
proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

server:
	go run ./cmd/

build: vet
		go build -o ./bin/api ./cmd/
