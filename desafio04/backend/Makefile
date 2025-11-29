.PHONY: help build run test clean docker-build docker-run wire

# Variáveis
SERVICE_NAME=weather-api
PORT=3000

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Desenvolvimento Local
build: ## Build da aplicação
	go build -o server ./cmd/server

run: ## Executa a aplicação localmente
	go run ./cmd/server

test: ## Executa testes
	go test -v ./...

test-coverage: ## Executa testes com coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

wire: ## Gera código de injeção de dependências
	cd cmd/server && wire

clean: ## Limpa arquivos gerados
	rm -f server
	rm -f coverage.out coverage.html
	rm -f cmd/server/wire_gen.go

tidy: ## Organiza dependências Go
	go mod tidy

fmt: ## Formata código Go
	go fmt ./...

lint: ## Executa linter
	golangci-lint run

dev: wire build run ## Workflow de desenvolvimento (wire + build + run)

# Docker
docker-build: ## Build da imagem Docker
	docker build -t $(SERVICE_NAME):latest .

docker-run: ## Executa container localmente
	docker run --rm -p $(PORT):8080 --env-file .env $(SERVICE_NAME):latest
