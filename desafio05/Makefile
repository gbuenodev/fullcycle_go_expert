.PHONY: run docker-up docker-down docker-clean docker-logs build help

run: ## Start MongoDB and run the application
	docker-compose up -d
	go run cmd/auction/main.go

docker-up: ## Start MongoDB container
	docker-compose up -d

docker-down: ## Stop MongoDB container
	docker-compose down

docker-clean: ## Stop MongoDB and remove volumes (fresh start)
	docker-compose down -v

docker-logs: ## Follow MongoDB logs
	docker-compose logs -f mongodb

build: ## Build the Go binary
	go build -o bin/auction cmd/auction/main.go

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
