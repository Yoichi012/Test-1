.PHONY: build run test clean docker-build docker-run help

# Variables
BINARY_NAME=shivu
MAIN_PATH=./cmd/shivu
DOCKER_IMAGE=shivu-bot

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)/main.go
	@echo "Build complete!"

run: ## Run the application
	@echo "Running $(BINARY_NAME)..."
	@go run $(MAIN_PATH)/main.go

dev: ## Run with live reload (requires air)
	@air

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@echo "Clean complete!"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies downloaded!"

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .
	@echo "Docker image built!"

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	@docker run -d --env-file .env --name $(BINARY_NAME) $(DOCKER_IMAGE)
	@echo "Docker container started!"

docker-stop: ## Stop Docker container
	@echo "Stopping Docker container..."
	@docker stop $(BINARY_NAME)
	@docker rm $(BINARY_NAME)
	@echo "Docker container stopped!"

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run ./...

format: ## Format code
	@echo "Formatting code..."
	@gofmt -s -w .
	@echo "Format complete!"

install: ## Install the application
	@echo "Installing $(BINARY_NAME)..."
	@go install $(MAIN_PATH)/main.go
	@echo "Installation complete!"