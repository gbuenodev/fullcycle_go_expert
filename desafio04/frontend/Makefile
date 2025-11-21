.PHONY: help install dev build preview lint clean docker-build docker-run

# Variáveis
SERVICE_NAME=weather-app
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
