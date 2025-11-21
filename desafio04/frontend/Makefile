.PHONY: help install dev build preview lint clean docker-build docker-run deploy update-env logs status delete

# Variáveis
SERVICE_NAME=weather-app
PROJECT_ID ?= $(shell gcloud config get-value project)
REGION=us-central1
BACKEND_URL ?= http://localhost:3000

help: ## Mostra esta mensagem de ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Desenvolvimento Local
install: ## Instala dependências
	npm install

dev: ## Inicia servidor de desenvolvimento
	npm run dev

build: ## Build de produção
	npm run build

preview: ## Preview do build de produção
	npm run preview

lint: ## Executa linter
	npm run lint

clean: ## Limpa node_modules e dist
	rm -rf node_modules dist

# Docker
docker-build: ## Build da imagem Docker
	docker build -t $(SERVICE_NAME) .

docker-run: ## Executa container localmente
	docker run -p 8080:8080 --env VITE_BACKEND_URL=$(BACKEND_URL) $(SERVICE_NAME)

docker-test: docker-build docker-run ## Build e executa container

# Google Cloud Run
setup-gcloud: ## Configura gcloud e habilita APIs
	@echo "Configurando gcloud..."
	gcloud config set project $(PROJECT_ID)
	gcloud config set run/region $(REGION)
	gcloud services enable cloudbuild.googleapis.com
	gcloud services enable run.googleapis.com
	@echo "Setup concluído!"

deploy: ## Deploy no Cloud Run
	@echo "Deploying $(SERVICE_NAME)..."
	@if [ -z "$(BACKEND_URL)" ]; then echo "Erro: BACKEND_URL não definido"; echo "Uso: make deploy BACKEND_URL=https://sua-url.run.app"; exit 1; fi
	gcloud run deploy $(SERVICE_NAME) \
		--source . \
		--platform managed \
		--region $(REGION) \
		--allow-unauthenticated \
		--port 8080 \
		--set-env-vars VITE_BACKEND_URL=$(BACKEND_URL)
	@echo "Deployed to Cloud Run!"

update-env: ## Atualiza variável de ambiente BACKEND_URL
	@echo "Atualizando variável de ambiente..."
	gcloud run services update $(SERVICE_NAME) \
		--update-env-vars VITE_BACKEND_URL=$(BACKEND_URL) \
		--region $(REGION)

logs: ## Mostra logs do serviço
	gcloud run services logs read $(SERVICE_NAME) --region $(REGION)

logs-tail: ## Mostra logs em tempo real
	gcloud run services logs tail $(SERVICE_NAME) --region $(REGION)

status: ## Mostra status do serviço
	gcloud run services describe $(SERVICE_NAME) --region $(REGION)

url: ## Mostra URL do serviço
	@gcloud run services describe $(SERVICE_NAME) --region $(REGION) --format="value(status.url)"

delete: ## Deleta o serviço do Cloud Run
	@echo "Deletando serviço $(SERVICE_NAME)..."
	gcloud run services delete $(SERVICE_NAME) --region $(REGION) --quiet

# Configurações Avançadas
set-scaling: ## Configura autoscaling (uso: make set-scaling MIN=0 MAX=10)
	gcloud run services update $(SERVICE_NAME) \
		--min-instances $(MIN) \
		--max-instances $(MAX) \
		--region $(REGION)

set-resources: ## Configura recursos (uso: make set-resources MEMORY=256Mi CPU=1)
	gcloud run services update $(SERVICE_NAME) \
		--memory $(MEMORY) \
		--cpu $(CPU) \
		--region $(REGION)

# Testes
health-check: ## Testa health check do serviço
	@URL=$$(gcloud run services describe $(SERVICE_NAME) --region $(REGION) --format="value(status.url)"); \
	curl -f $$URL/health || echo "Health check falhou"

# CI/CD
ci-build: install lint build ## Pipeline CI: instala, lint e build

# Informações
info: ## Mostra informações do projeto
	@echo "Project ID: $(PROJECT_ID)"
	@echo "Region: $(REGION)"
	@echo "Service Name: $(SERVICE_NAME)"
	@echo "Backend URL: $(BACKEND_URL)"
