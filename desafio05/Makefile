.PHONY: help run dev build test clean docker-up docker-down docker-build docker-logs docker-clean

.DEFAULT_GOAL := help

help: ## Display available commands
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Start everything with Docker (MongoDB + API)
	docker-compose up -d

dev: ## Run locally (start MongoDB only, run Go app locally)
	docker-compose up -d mongodb
	go run cmd/auction/main.go

build: ## Build the Go binary
	go build -o bin/auction cmd/auction/main.go

test: ## Run all tests
	go test -v ./...

clean: ## Remove everything (containers, volumes, binaries)
	docker-compose down -v
	rm -rf bin/

docker-up: ## Start services in detached mode
	docker-compose up -d

docker-down: ## Stop all services
	docker-compose down

docker-build: ## Rebuild images without cache
	docker-compose build --no-cache

docker-logs: ## Follow API logs
	docker-compose logs -f auction-api

docker-clean: ## Stop services and remove volumes
	docker-compose down -v
