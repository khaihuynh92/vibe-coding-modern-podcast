# Makefile for Podsite Backend (Go)

.PHONY: build run dev test clean fmt lint vet deps docker-build docker-run help

# Variables
BINARY_NAME=podsite-backend
MAIN_PATH=./cmd/server
BUILD_DIR=./bin
DOCKER_IMAGE=podsite-backend

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Development
dev: ## Run the application in development mode with auto-reload
	@echo "Starting development server..."
	@go run $(MAIN_PATH)/main.go

run: build ## Build and run the application
	@echo "Running $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME)

# Building
build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

build-linux: ## Build for Linux
	@echo "Building $(BINARY_NAME) for Linux..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux $(MAIN_PATH)

# Testing
test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

test-race: ## Run tests with race detection
	@echo "Running tests with race detection..."
	@go test -race -v ./...

test-handlers: ## Run handler tests only
	@echo "Running handler tests..."
	@go test -v ./internal/handlers/...

test-models: ## Run model tests only
	@echo "Running model tests..."
	@go test -v ./internal/models/...

test-middleware: ## Run middleware tests only
	@echo "Running middleware tests..."
	@go test -v ./internal/middleware/...

benchmark: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

# Code quality
fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

# Dependencies
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download

deps-update: ## Update dependencies
	@echo "Updating dependencies..."
	@go get -u ./...
	@go mod tidy

deps-verify: ## Verify dependencies
	@echo "Verifying dependencies..."
	@go mod verify

# Cleanup
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

# Docker
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	@docker run -p 3001:3001 --env-file .env $(DOCKER_IMAGE)

docker-dev: ## Run Docker container in development mode
	@echo "Running Docker container in development mode..."
	@docker run -p 3001:3001 -v $(PWD):/app --env-file .env $(DOCKER_IMAGE) make dev

# Docker Compose
docker-compose-up: ## Start services with Docker Compose
	@echo "Starting services with Docker Compose..."
	@docker-compose up -d

docker-compose-down: ## Stop services with Docker Compose
	@echo "Stopping services with Docker Compose..."
	@docker-compose down

docker-compose-build: ## Build services with Docker Compose
	@echo "Building services with Docker Compose..."
	@docker-compose build

docker-compose-logs: ## View logs with Docker Compose
	@echo "Viewing logs with Docker Compose..."
	@docker-compose logs -f

docker-compose-restart: ## Restart services with Docker Compose
	@echo "Restarting services with Docker Compose..."
	@docker-compose restart

# Health check
health: ## Check if the server is running
	@echo "Checking server health..."
	@curl -f http://localhost:3001/health || echo "Server is not running"

# Generate
generate: ## Run go generate
	@echo "Running go generate..."
	@go generate ./...

# Install tools
install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/swaggo/swag/cmd/swag@latest

# Swagger
swagger: ## Generate Swagger documentation
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/server/main.go -o ./docs
