.PHONY: help build run test clean docker-build docker-run wire deploy

# Variables
APP_NAME=weather-api
DOCKER_IMAGE=gcr.io/$(GCP_PROJECT_ID)/$(APP_NAME)
PORT=3000

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the application binary
	@echo "Building application..."
	go build -o server ./cmd/server
	@echo "✓ Build complete!"

run: ## Run the application locally
	@echo "Running application on port $(PORT)..."
	go run ./cmd/server

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report generated: coverage.html"

wire: ## Generate wire dependency injection
	@echo "Generating wire code..."
	cd cmd/server && wire
	@echo "✓ Wire generation complete!"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f server
	rm -f coverage.out coverage.html
	rm -f cmd/server/wire_gen.go
	@echo "✓ Clean complete!"

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):latest .
	@echo "✓ Docker image built!"

docker-run: ## Run Docker container locally
	@echo "Running Docker container on port $(PORT)..."
	docker run --rm -p $(PORT):8080 --env-file .env $(APP_NAME):latest

docker-push: ## Push Docker image to GCR
	@echo "Tagging and pushing to GCR..."
	docker tag $(APP_NAME):latest $(DOCKER_IMAGE):latest
	docker push $(DOCKER_IMAGE):latest
	@echo "✓ Image pushed to GCR!"

deploy: ## Deploy to Cloud Run
	@echo "Deploying to Cloud Run..."
	gcloud run deploy $(APP_NAME) \
		--image $(DOCKER_IMAGE):latest \
		--platform managed \
		--region us-central1 \
		--allow-unauthenticated \
		--set-env-vars PORT=8080,WEATHER_API_KEY=$(WEATHER_API_KEY),WEATHER_API_BASE_URL=$(WEATHER_API_BASE_URL),VIACEP_BASE_URL=$(VIACEP_BASE_URL)
	@echo "✓ Deployed to Cloud Run!"

deploy-dev: docker-build docker-push deploy ## Build, push and deploy in one command

tidy: ## Tidy go modules
	@echo "Tidying go modules..."
	go mod tidy
	@echo "✓ Modules tidied!"

fmt: ## Format Go code
	@echo "Formatting code..."
	go fmt ./...
	@echo "✓ Code formatted!"

lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run
	@echo "✓ Lint complete!"

dev: wire build run ## Generate wire, build and run (development workflow)
