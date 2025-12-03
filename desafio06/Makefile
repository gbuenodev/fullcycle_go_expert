.PHONY: help build up down restart logs redis-cli test-ip test-token test-auth clean

# Variables
DOCKER_COMPOSE = docker-compose
APP_NAME = rate-limiter 

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build Docker images
	$(DOCKER_COMPOSE) build

up: ## Start all services
	$(DOCKER_COMPOSE) up -d
	@echo "✅ Services started"
	@echo "App: http://localhost:8080"
	@echo "Redis: localhost:6379"

down: ## Stop all services
	$(DOCKER_COMPOSE) down
	@echo "✅ Services stopped"

restart: down up ## Restart all services

logs: ## Show application logs
	$(DOCKER_COMPOSE) logs -f ${APP_NAME}

logs-redis: ## Show Redis logs
	$(DOCKER_COMPOSE) logs -f redis

redis-cli: ## Access Redis CLI
	docker exec -it $$(docker ps -qf "name=redis") redis-cli

redis-keys: ## Show all rate limit keys in Redis
	docker exec -it $$(docker ps -qf "name=redis") redis-cli KEYS "ratelimit:*"

redis-flush: ## Flush all Redis data (CAREFUL!)
	docker exec -it $$(docker ps -qf "name=redis") redis-cli FLUSHALL
	@echo "⚠️  Redis flushed"

test-auth: ## Get a test token
	@echo "Getting basic token:"
	@curl -s "http://localhost:8080/auth?tier=basic" | jq
	@echo ""
	@echo "Getting premium token:"
	@curl -s "http://localhost:8080/auth?tier=premium" | jq

test-ip: ## Test IP-based rate limiting (sends 15 requests)
	@echo "Sending 15 requests (limit is 10)..."
	@for i in $$(seq 1 15); do \
		status=$$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/); \
		echo "Request $$i - Status: $$status"; \
	done

test-token: ## Test token-based rate limiting (requires TOKEN env var)
	@if [ -z "$(TOKEN)" ]; then \
		echo "❌ Error: TOKEN variable required"; \
		echo "Usage: make test-token TOKEN=your-token-here"; \
		exit 1; \
	fi
	@echo "Testing with token: $(TOKEN)"
	@echo "Sending 55 requests (basic=50, premium=200)..."
	@for i in $$(seq 1 55); do \
		status=$$(curl -s -o /dev/null -w "%{http_code}" -H "API_KEY: $(TOKEN)" http://localhost:8080/); \
		echo "Request $$i - Status: $$status"; \
	done

test-concurrent: ## Test concurrent requests
	@echo "Sending 20 concurrent requests..."
	@seq 1 20 | xargs -P 20 -I {} curl -s -o /dev/null -w "Request {} - Status: %{http_code}\n" http://localhost:8080/

status: ## Show service status
	$(DOCKER_COMPOSE) ps

clean: down ## Clean up everything
	$(DOCKER_COMPOSE) down -v
	docker system prune -f
	@echo "✅ Cleanup complete"

dev-up: ## Start only Redis for local development
	$(DOCKER_COMPOSE) up -d redis
	@echo "✅ Redis started for development"
	@echo "Run app locally: go run cmd/server/main.go"

dev-down: ## Stop development Redis
	$(DOCKER_COMPOSE) stop redis

run-local: ## Run application locally (requires Redis running)
	@echo "Starting application locally..."
	REDIS_ADDR=localhost:6379 go run cmd/server/main.go

watch: ## Watch logs in real-time
	$(DOCKER_COMPOSE) logs -f

health: ## Check service health
	@echo "Checking app health..."
	@curl -s http://localhost:8080/ && echo "✅ App is healthy" || echo "❌ App is down"
	@echo "Checking Redis..."
	@docker exec $$(docker ps -qf "name=redis") redis-cli PING && echo "✅ Redis is healthy" || echo "❌ Redis is down"
