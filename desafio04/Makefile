.PHONY: help up down restart logs build clean ps health test quickstart

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'

# ===== Docker Compose =====
up: ## Sobe todos os serviços (backend + frontend + input-service + observability)
	docker-compose up -d
	@sleep 5
	@make health

down: ## Para e remove todos os containers
	docker-compose down

restart: ## Reinicia todos os serviços
	docker-compose restart

stop: ## Para os serviços sem remover
	docker-compose stop

start: ## Inicia os serviços parados
	docker-compose start

build: ## Build das imagens
	docker-compose build

build-no-cache: ## Build das imagens sem cache
	docker-compose build --no-cache

rebuild: down build up ## Rebuild completo (down + build + up)

quickstart: build up ## Quick start: build e up

# ===== Logs =====
logs: ## Mostra logs de todos os serviços
	docker-compose logs -f

logs-frontend: ## Mostra logs do frontend
	docker-compose logs -f frontend

logs-backend: ## Mostra logs do backend
	docker-compose logs -f backend

logs-input: ## Mostra logs do input-service
	docker-compose logs -f input-service

logs-zipkin: ## Mostra logs do Zipkin
	docker-compose logs -f zipkin

logs-prometheus: ## Mostra logs do Prometheus
	docker-compose logs -f prometheus

# ===== Status e Monitoramento =====
ps: ## Mostra status dos containers
	docker-compose ps

health: ## Verifica health dos serviços
	@curl -sf http://localhost:3000/health > /dev/null && echo "Backend: OK" || echo "Backend: FAILED"
	@curl -sf http://localhost:8080/health > /dev/null && echo "Frontend: OK" || echo "Frontend: FAILED"
	@curl -sf http://localhost:9411/health > /dev/null && echo "Zipkin: OK" || echo "Zipkin: FAILED"
	@curl -sf http://localhost:9090/-/healthy > /dev/null && echo "Prometheus: OK" || echo "Prometheus: FAILED"

# ===== Testes =====
test-backend-post: ## Testa POST /weather do backend
	@curl -X POST http://localhost:3000/weather \
		-H "Content-Type: application/json" \
		-d '{"cep":"01310100"}'

test-input-service: ## Testa input-service
	@curl -X POST http://localhost:8000/weather \
		-H "Content-Type: application/json" \
		-d '{"cep":"01310100"}'

test-distributed-tracing: test-input-service ## Testa distributed tracing (input-service → backend)
	@echo "Trace: http://localhost:9411"

test-all: health test-backend-post test-input-service ## Executa todos os testes

# ===== Observabilidade =====
zipkin: ## Abre Zipkin UI no browser
	@open http://localhost:9411 2>/dev/null || xdg-open http://localhost:9411 2>/dev/null || echo "http://localhost:9411"

prometheus: ## Abre Prometheus UI no browser
	@open http://localhost:9090 2>/dev/null || xdg-open http://localhost:9090 2>/dev/null || echo "http://localhost:9090"

metrics: ## Mostra métricas do backend
	@curl -s http://localhost:3000/metrics | grep -E "^(http_|go_)" | head -20

# ===== Limpeza =====
clean: down ## Para serviços e remove volumes
	docker-compose down -v

clean-all: ## Remove tudo (containers, volumes, imagens)
	docker-compose down -v --rmi all

prune: ## Remove containers, networks e images não utilizados
	docker system prune -f

# ===== Desenvolvimento =====
dev-frontend: ## Entra no container do frontend
	docker-compose exec frontend sh

dev-backend: ## Entra no container do backend
	docker-compose exec backend sh

dev-input: ## Entra no container do input-service
	docker-compose exec input-service sh

# ===== Informações =====
info: ## Mostra informações do ambiente
	@docker-compose ps
	@echo ""
	@echo "Frontend:      http://localhost:8080"
	@echo "Backend:       http://localhost:3000"
	@echo "Input Service: http://localhost:8000"
	@echo "Zipkin:        http://localhost:9411"
	@echo "Prometheus:    http://localhost:9090"
