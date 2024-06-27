.DEFAULT_GOAL := build

# Directories for each service
API_GATEWAY_DIR := api-gateway
AUTH_DIR := auth
USER_DIR := user
NOTIFICATION_DIR := notification
TASK_DIR := task

# Targets for vetting, generating protobuf code, building, and running services
.PHONY: vet proto build run

vet:
	cd $(API_GATEWAY_DIR) && go vet ./...
	cd $(AUTH_DIR) && go vet ./...
	cd $(USER_DIR) && go vet ./...
	cd $(NOTIFICATION_DIR) && go vet ./...
	cd $(TASK_DIR) && go vet ./...

proto:
	protoc ./**/pb/*.proto --go_out=. --go-grpc_out=.

build: vet
	@echo "Building API Gateway..."
	cd $(API_GATEWAY_DIR) && go build -o bin/api ./cmd/main.go
	@echo "Building Auth Service..."
	cd $(AUTH_DIR) && go build -o bin/auth ./cmd/main.go
	@echo "Building User Service..."
	cd $(USER_DIR) && go build -o bin/user ./cmd/main.go
	@echo "Building Notification Service..."
	cd $(NOTIFICATION_DIR) && go build -o bin/notification ./cmd/main.go
	@echo "Building Task Service..."
	cd $(TASK_DIR) && go build -o bin/task ./cmd/main.go

run: build
	@echo "Running API Gateway..."
	cd $(API_GATEWAY_DIR) && go run ./cmd/main.go
	@echo "Running Auth Service..."
	cd $(AUTH_DIR) && go run ./cmd/main.go
	@echo "Running User Service..."
	cd $(USER_DIR) && go run ./cmd/main.go
	@echo "Running Notification Service..."
	cd $(NOTIFICATION_DIR) && go run ./cmd/main.go
	@echo "Running Task Service..."
	cd $(TASK_DIR) && go run ./cmd/main.go

clean:
	@echo "Cleaning API Gateway..."
	cd $(API_GATEWAY_DIR) && go clean
	@echo "Cleaning Auth Service..."
	cd $(AUTH_DIR) && go clean
	@echo "Cleaning User Service..."
	cd $(USER_DIR) && go clean
	@echo "Cleaning Notification Service..."
	cd $(NOTIFICATION_DIR) && go clean
	@echo "Cleaning Task Service..."
	cd $(TASK_DIR) && go clean

