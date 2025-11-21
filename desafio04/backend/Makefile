.PHONY: help build run test clean docker-build docker-run wire deploy

# Variáveis
SERVICE_NAME=weather-api
PORT=3000
REGION=us-central1

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Desenvolvimento Local
build: ## Build da aplicação
	@echo "Building application..."
	go build -o server ./cmd/server
	@echo "Build complete!"

run: ## Executa a aplicação localmente
	@echo "Running application on port $(PORT)..."
	go run ./cmd/server

test: ## Executa testes
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Executa testes com coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

wire: ## Gera código de injeção de dependências
	@echo "Generating wire code..."
	cd cmd/server && wire
	@echo "Wire generation complete!"

clean: ## Limpa arquivos gerados
	@echo "Cleaning..."
	rm -f server
	rm -f coverage.out coverage.html
	rm -f cmd/server/wire_gen.go
	@echo "Clean complete!"

tidy: ## Organiza dependências Go
	@echo "Tidying go modules..."
	go mod tidy
	@echo "Modules tidied!"

fmt: ## Formata código Go
	@echo "Formatting code..."
	go fmt ./...
	@echo "Code formatted!"

lint: ## Executa linter
	@echo "Running linter..."
	golangci-lint run
	@echo "Lint complete!"

dev: wire build run ## Workflow de desenvolvimento (wire + build + run)

# Docker
docker-build: ## Build da imagem Docker
	@echo "Building Docker image..."
	docker build -t $(SERVICE_NAME):latest .
	@echo "Docker image built!"

docker-run: ## Executa container localmente
	@echo "Running Docker container on port $(PORT)..."
	docker run --rm -p $(PORT):8080 --env-file .env $(SERVICE_NAME):latest

# Google Cloud Run
deploy: ## Deploy no Cloud Run
	@echo "Deploying to Cloud Run..."
	@if [ -z "$(WEATHER_API_KEY)" ]; then echo "Erro: WEATHER_API_KEY não definido"; exit 1; fi
	gcloud run deploy $(SERVICE_NAME) \
		--source . \
		--platform managed \
		--region $(REGION) \
		--allow-unauthenticated \
		--port 8080 \
		--set-env-vars WEATHER_API_KEY=$(WEATHER_API_KEY)
	@echo "Deployed to Cloud Run!"
