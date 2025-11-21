.PHONY: help up down restart logs logs-frontend logs-backend build clean ps health test-backend test-frontend test-all

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Docker Compose
up: ## Sobe todos os serviços (backend + frontend)
	@echo "Subindo serviços..."
	docker-compose up -d
	@echo "Aguardando serviços ficarem prontos..."
	@sleep 5
	@make health
	@echo ""
	@echo "Serviços rodando!"
	@echo "Frontend: http://localhost:8080"
	@echo "Backend: http://localhost:3000"

down: ## Para e remove todos os containers
	@echo "Parando serviços..."
	docker-compose down

restart: ## Reinicia todos os serviços
	@echo "Reiniciando serviços..."
	docker-compose restart

stop: ## Para os serviços sem remover
	@echo "Parando serviços..."
	docker-compose stop

start: ## Inicia os serviços parados
	@echo "Iniciando serviços..."
	docker-compose start

build: ## Build das imagens
	@echo "Building imagens..."
	docker-compose build

build-no-cache: ## Build das imagens sem cache
	@echo "Building imagens sem cache..."
	docker-compose build --no-cache

rebuild: down build up ## Rebuild completo (down + build + up)

quickstart: build up ## Quick start: build e up
	@echo ""
	@echo "Ambiente pronto!"
	@echo ""
	@echo "Acesse:"
	@echo "  Frontend: http://localhost:8080"
	@echo "  Backend:  http://localhost:3000"
	@echo ""
	@echo "Para ver logs: make logs"
	@echo "Para parar: make down"

# Logs
logs: ## Mostra logs de todos os serviços
	docker-compose logs -f

logs-frontend: ## Mostra logs apenas do frontend
	docker-compose logs -f frontend

logs-backend: ## Mostra logs apenas do backend
	docker-compose logs -f backend

# Status e Testes
ps: ## Mostra status dos containers
	docker-compose ps

health: ## Verifica health dos serviços
	@echo "Verificando health dos serviços..."
	@echo ""
	@echo "Backend:"
	@curl -f http://localhost:3000/health 2>/dev/null && echo "  OK" || echo "  Falhou"
	@echo ""
	@echo "Frontend:"
	@curl -f http://localhost:8080/health 2>/dev/null && echo "  OK" || echo "  Falhou"
	@echo ""

test-backend: ## Testa endpoint do backend
	@echo "Testando backend..."
	@echo "GET /health"
	curl -v http://localhost:3000/health
	@echo ""

test-frontend: ## Testa frontend
	@echo "Testando frontend..."
	@echo "GET /health"
	curl -v http://localhost:8080/health
	@echo ""

test-api: ## Testa API com CEP de exemplo
	@echo "Testando API com CEP 01310-100..."
	curl -v http://localhost:3000/weather/01310100
	@echo ""

test-all: health test-backend test-frontend ## Executa todos os testes

# Limpeza
clean: down ## Para serviços e remove volumes
	@echo "Limpando volumes..."
	docker-compose down -v

clean-all: ## Remove tudo (containers, volumes, imagens)
	@echo "Removendo tudo..."
	docker-compose down -v --rmi all

prune: ## Remove containers, networks e images não utilizados
	@echo "Limpando sistema Docker..."
	docker system prune -f

# Desenvolvimento
dev-frontend: ## Entra no container do frontend
	docker-compose exec frontend sh

dev-backend: ## Entra no container do backend
	docker-compose exec backend sh

# Informações
info: ## Mostra informações do ambiente
	@echo "Informações do Ambiente:"
	@echo ""
	@echo "Docker Compose:"
	docker-compose version
	@echo ""
	@echo "Containers:"
	@docker-compose ps
	@echo ""
	@echo "URLs:"
	@echo "  Frontend: http://localhost:8080"
	@echo "  Backend:  http://localhost:3000"
	@echo ""

# Google Cloud Run
setup-gcloud: ## Configura Google Cloud
	@echo "Configurando Google Cloud..."
	@if [ -z "$(GCP_PROJECT_ID)" ]; then echo "Erro: GCP_PROJECT_ID não definido"; echo "Uso: make setup-gcloud GCP_PROJECT_ID=seu-projeto"; exit 1; fi
	gcloud config set project $(GCP_PROJECT_ID)
	gcloud services enable run.googleapis.com
	gcloud services enable cloudbuild.googleapis.com
	@echo "Google Cloud configurado!"

deploy-backend: ## Deploy do backend no Cloud Run
	@echo "Deploying backend to Cloud Run..."
	@if [ -z "$(WEATHER_API_KEY)" ]; then echo "Erro: WEATHER_API_KEY não definido"; exit 1; fi
	cd backend && make deploy WEATHER_API_KEY=$(WEATHER_API_KEY)
	@echo ""
	@echo "Backend deployed!"
	@echo "Copie a URL exibida acima para usar no deploy do frontend"

deploy-frontend: ## Deploy do frontend no Cloud Run (requer BACKEND_URL)
	@echo "Deploying frontend to Cloud Run..."
	@if [ -z "$(BACKEND_URL)" ]; then echo "Erro: BACKEND_URL não definido"; exit 1; fi
	cd frontend && make deploy BACKEND_URL=$(BACKEND_URL)
	@echo ""
	@echo "Frontend deployed!"

deploy-all: deploy-backend ## Deploy completo (backend + frontend)
	@echo ""
	@echo "Aguarde o backend ficar disponível..."
	@echo "Copie a URL do backend e execute:"
	@echo ""
	@echo "  make deploy-frontend BACKEND_URL=https://sua-url-backend.run.app"
	@echo ""
