# Weather App - Consulta de Clima por CEP

Aplicação fullstack para consultar informações de clima baseado em CEP brasileiro.

## Arquitetura

```
desafio04/
├── backend/          # API Go com Clean Architecture
├── frontend/         # React + TypeScript + MUI
├── docker-compose.yml
└── Makefile
```

## Stack

### Backend
- Go 1.25
- Clean Architecture (Entity, UseCase, Gateway, Infra)
- Dependency Injection com Wire
- Router: chi
- Configuration: Viper
- APIs Externas: ViaCEP + WeatherAPI

### Frontend
- React 18
- TypeScript
- Material-UI
- Vite
- Axios
- Zod

## Quick Start

### Pré-requisitos

- Docker e Docker Compose
- Make (opcional, mas recomendado)

### Subir ambiente completo

```bash
# Ver comandos disponíveis
make help

# Build e subir tudo
make quickstart

# Ou manualmente
docker-compose up -d
```

Acesse:
- **Frontend**: http://localhost:8080
- **Backend**: http://localhost:3000

### Parar ambiente

```bash
make down
```

## Comandos Principais

```bash
# Gerenciamento
make up              # Sobe todos os serviços
make down            # Para e remove containers
make restart         # Reinicia serviços
make logs            # Ver logs de todos
make logs-frontend   # Ver logs do frontend
make logs-backend    # Ver logs do backend

# Build
make build           # Build das imagens
make rebuild         # Rebuild completo

# Monitoramento
make ps              # Status dos containers
make health          # Verifica health dos serviços
make test-all        # Testa todos os endpoints

# Desenvolvimento
make dev-frontend    # Shell no container frontend
make dev-backend     # Shell no container backend

# Limpeza
make clean           # Remove containers e volumes
make clean-all       # Remove tudo (imagens também)
```

## Desenvolvimento

### Frontend

```bash
cd frontend/
make help            # Ver comandos do frontend
make dev             # Desenvolvimento local
```

Ver [frontend/README.md](frontend/README.md) para mais detalhes.

### Backend

```bash
cd backend/
# Comandos do backend
```

Ver [backend/README.md](backend/README.md) para mais detalhes.

## Docker Compose

### Serviços

- **backend**: API Go na porta 3000
- **frontend**: React App na porta 8080
- **network**: `weather-network` para comunicação entre serviços

### Health Checks

Ambos os serviços possuem health checks configurados:
- Frontend: `GET http://localhost:8080/health`
- Backend: `GET http://localhost:3000/health`

## API

### Endpoint Principal

```
GET /weather/{cep}
```

**Parâmetros:**
- `cep`: CEP brasileiro (8 dígitos, apenas números)

**Respostas:**

**200 OK**
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

**404 Not Found**
```json
{
  "message": "can not find zipcode"
}
```

**422 Unprocessable Entity**
```json
{
  "message": "invalid zipcode"
}
```

## Testes

### Testar Health Checks

```bash
# Todos
make health

# Individual
curl http://localhost:8080/health  # Frontend
curl http://localhost:3000/health  # Backend
```

### Testar API

```bash
# Usando make
make test-api

# Ou curl direto
curl http://localhost:3000/weather/01310100
```

## Deploy no Cloud Run

### Pré-requisitos

1. **Google Cloud SDK** instalado e configurado:
```bash
gcloud auth login
gcloud config set project SEU_PROJECT_ID
```

2. **Habilitar APIs** necessárias:
```bash
make setup-gcloud GCP_PROJECT_ID=seu-projeto

# Ou manualmente
gcloud services enable run.googleapis.com
gcloud services enable cloudbuild.googleapis.com
```

3. **Obter API Key** do WeatherAPI:
   - Acesse: https://www.weatherapi.com/
   - Crie uma conta gratuita
   - Copie sua API Key

### Deploy Completo (Backend + Frontend)

**Opção 1: Deploy com comandos da raiz**

```bash
# 1. Configurar variáveis
export GCP_PROJECT_ID=seu-projeto-gcp
export WEATHER_API_KEY=sua-chave-weatherapi

# 2. Deploy do backend
make deploy-backend

# 3. Copiar URL do backend (será exibida após o deploy)
export BACKEND_URL=https://weather-api-xxx.run.app

# 4. Deploy do frontend
make deploy-frontend BACKEND_URL=$BACKEND_URL
```

**Opção 2: Deploy individual**

#### Backend
```bash
cd backend/
export WEATHER_API_KEY=sua-chave
make deploy WEATHER_API_KEY=$WEATHER_API_KEY
```

URL do backend será exibida após deploy (ex: `https://weather-api-xxx.run.app`)

#### Frontend
```bash
cd frontend/
make deploy BACKEND_URL=https://weather-api-xxx.run.app
```

URL do frontend será exibida após deploy (ex: `https://weather-app-xxx.run.app`)

### Configuração de Variáveis de Ambiente

#### Backend (Cloud Run)
```bash
gcloud run services update weather-api \
  --set-env-vars="PORT=8080,WEATHER_API_KEY=sua-chave,WEATHER_API_BASE_URL=https://api.weatherapi.com/v1,VIACEP_BASE_URL=https://viacep.com.br/ws"
```

#### Frontend (Build Time)
```bash
# Durante o deploy, passe a URL do backend
make deploy-source BACKEND_URL=https://weather-api-xxx.run.app
```

### Verificação

```bash
# Testar backend
curl https://weather-api-xxx.run.app/weather/01310100

# Acessar frontend no browser
open https://weather-app-xxx.run.app
```

## Troubleshooting

### Porta já em uso

```bash
# Verificar o que está usando a porta
lsof -i :8080  # Frontend
lsof -i :3000  # Backend

# Ou parar tudo
make down
```

### Container não sobe

```bash
# Ver logs
make logs

# Rebuild sem cache
make build-no-cache
make up
```

### Backend não responde

```bash
# Verificar status
make ps

# Ver logs do backend
make logs-backend

# Verificar health
curl http://localhost:3000/health
```

### Frontend não conecta ao backend

1. Verificar se backend está rodando: `make health`
2. Verificar network: `docker network inspect desafio04_weather-network`
3. Verificar variável de ambiente: `docker-compose exec frontend env | grep BACKEND`

## Estrutura de URLs

### Desenvolvimento Local (Docker Compose)
- Frontend: http://localhost:8080
- Backend: http://localhost:3000
- Frontend chama backend via: http://backend:3000 (internal)

### Produção (Cloud Run)
- Frontend: https://weather-app-xxx.run.app
- Backend: https://weather-api-xxx.run.app
- Frontend chama backend via: VITE_BACKEND_URL (env var)

## Variáveis de Ambiente

### Frontend
- `VITE_BACKEND_URL`: URL do backend (build time)

### Backend
- `PORT`: Porta do servidor (default: 3000)

## Licença

MIT
