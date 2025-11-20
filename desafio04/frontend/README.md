# Consulta Clima por CEP - Frontend

Frontend React + TypeScript para consultar informações de clima baseado no CEP brasileiro.

## Tecnologias

- **React 18** - Biblioteca UI
- **TypeScript** - Tipagem estática
- **Vite** - Build tool e dev server
- **Material-UI (MUI)** - Componentes UI
- **Axios** - Cliente HTTP
- **Zod** - Validação de esquemas
- **React Input Mask** - Máscara de input

## Arquitetura

```
src/
├── api/              # Configuração Axios
├── components/       # Componentes React
│   ├── Loading.tsx
│   ├── WeatherForm.tsx
│   └── WeatherResult.tsx
├── context/          # Context API
│   ├── WeatherContext.tsx
│   └── WeatherProvider.tsx
├── hooks/            # Custom hooks
│   └── useWeather.ts
├── services/         # Serviços de API
│   └── weatherService.ts
└── utils/            # Utilitários
    └── validation.ts
```

## Funcionalidades

- Formulário com validação de CEP brasileiro
- Máscara automática (formato: 00000-000)
- Consulta de temperatura em Celsius, Fahrenheit e Kelvin
- Tratamento de erros (404, 422)
- Loading states
- UI responsiva e moderna

## Makefile - Comandos Rápidos

Este projeto inclui um Makefile para facilitar operações comuns:

```bash
# Ver todos os comandos disponíveis
make help

# Desenvolvimento
make install          # Instalar dependências
make dev             # Servidor de desenvolvimento
make build           # Build de produção
make lint            # Executar linter

# Docker
make docker-build    # Build da imagem Docker
make docker-run      # Executar container localmente
make docker-test     # Build e executar

# Cloud Run - Deploy
make setup-gcloud    # Configurar gcloud e APIs
make deploy-source   # Deploy direto do código (recomendado)
make deploy          # Build imagem e deploy

# Cloud Run - Gerenciamento
make update-env BACKEND_URL=https://backend.run.app  # Atualizar variável
make logs            # Ver logs
make logs-tail       # Ver logs em tempo real
make status          # Ver status do serviço
make url             # Ver URL do serviço
make delete          # Deletar serviço

# Configurações
make set-scaling MIN=0 MAX=10               # Configurar autoscaling
make set-resources MEMORY=256Mi CPU=1       # Configurar recursos
```

### Variáveis do Makefile

Você pode sobrescrever as variáveis padrão:

```bash
# Exemplo: deploy com projeto e região específicos
make deploy PROJECT_ID=meu-projeto REGION=us-east1 BACKEND_URL=https://api.example.com
```

## Desenvolvimento Local

### Pré-requisitos

- Node.js 18+
- npm ou yarn

### Instalação

```bash
# Instalar dependências
npm install

# Configurar variáveis de ambiente
cp .env.example .env
# Editar .env com a URL do backend
```

### Executar

```bash
# Modo desenvolvimento
npm run dev

# Build de produção
npm run build

# Preview do build
npm run preview

# Lint
npm run lint
```

## Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
VITE_BACKEND_URL=http://localhost:3000
```

## Docker

### Build da imagem

```bash
docker build -t weather-frontend .
```

### Executar localmente

```bash
docker run -p 8080:8080 weather-frontend
```

Acesse: http://localhost:8080

## Deploy no Google Cloud Run

### Pré-requisitos

- Google Cloud CLI instalado
- Projeto GCP criado
- Billing habilitado

### Deploy Rápido com Makefile (Recomendado)

```bash
# 1. Configurar projeto e habilitar APIs (uma vez)
make setup-gcloud PROJECT_ID=seu-projeto-id

# 2. Deploy direto do código-fonte
make deploy-source BACKEND_URL=https://seu-backend.run.app

# 3. Ver URL do serviço
make url

# Atualizar variável de ambiente depois
make update-env BACKEND_URL=https://novo-backend.run.app
```

### Deploy Manual (Alternativa)

<details>
<summary>Clique para ver comandos gcloud diretos</summary>

1. **Autenticar no GCP**
```bash
gcloud auth login
gcloud config set project [PROJECT_ID]
```

2. **Habilitar APIs**
```bash
gcloud services enable cloudbuild.googleapis.com
gcloud services enable run.googleapis.com
gcloud services enable containerregistry.googleapis.com
```

3. **Deploy direto do código (Opção A - Recomendado)**
```bash
gcloud run deploy weather-frontend \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080 \
  --set-env-vars VITE_BACKEND_URL=[URL_DO_BACKEND]
```

4. **Ou Build e Deploy da imagem (Opção B)**
```bash
# Build e push
gcloud builds submit --tag gcr.io/[PROJECT_ID]/weather-frontend

# Deploy
gcloud run deploy weather-frontend \
  --image gcr.io/[PROJECT_ID]/weather-frontend \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080 \
  --set-env-vars VITE_BACKEND_URL=[URL_DO_BACKEND]
```

5. **Atualizar variável de ambiente**
```bash
gcloud run services update weather-frontend \
  --update-env-vars VITE_BACKEND_URL=https://backend-url.run.app \
  --region us-central1
```

</details>

### Deploy automático com GitHub Actions

Veja `.github/workflows/deploy.yml` para CI/CD automático.

## API Backend

O frontend espera um backend com o seguinte endpoint:

```
GET /weather/{cep}
```

**Respostas:**

- `200 OK`:
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

- `404 Not Found`: CEP não encontrado
- `422 Unprocessable Entity`: CEP inválido

## Scripts NPM

- `npm run dev` - Inicia servidor de desenvolvimento
- `npm run build` - Build de produção
- `npm run preview` - Preview do build de produção
- `npm run lint` - Executa ESLint

## Exemplos de Uso

### Desenvolvimento Local
```bash
# Setup inicial
make install
make dev
```

### Teste com Docker
```bash
# Build e teste local
make docker-build
make docker-run

# Ou tudo junto
make docker-test
```

### Deploy no Cloud Run
```bash
# Primeira vez
make setup-gcloud PROJECT_ID=meu-projeto
make deploy-source BACKEND_URL=https://api.example.com

# Deploys futuros
make deploy-source BACKEND_URL=https://api.example.com

# Monitoramento
make logs-tail    # Ver logs em tempo real
make status       # Ver status
make url          # Ver URL pública
```

### Atualizar após deploy do backend
```bash
# Atualizar URL do backend
make update-env BACKEND_URL=https://novo-backend-url.run.app

# Verificar se está funcionando
make health-check
```

## Troubleshooting

### Comando make não encontrado
```bash
# macOS
xcode-select --install

# Linux
sudo apt-get install build-essential  # Ubuntu/Debian
sudo yum groupinstall "Development Tools"  # CentOS/RHEL
```

### Erro de permissões no gcloud
```bash
gcloud auth login
gcloud auth application-default login
```

### Build falha no Cloud Run
```bash
# Ver logs do build
make logs

# Testar build localmente
make docker-build
```

## Licença

MIT
