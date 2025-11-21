# Guia de Deploy - Google Cloud Run

## Pré-requisitos

### 1. Google Cloud SDK

Instale o gcloud CLI:
```bash
# macOS
brew install --cask google-cloud-sdk

# Linux
curl https://sdk.cloud.google.com | bash

# Windows
# Baixe o instalador: https://cloud.google.com/sdk/docs/install
```

### 2. Autenticação e Projeto

```bash
# Login
gcloud auth login

# Criar ou selecionar projeto
gcloud projects create SEU-PROJETO-ID  # Ou use um existente
gcloud config set project SEU-PROJETO-ID

# Verificar projeto atual
gcloud config get-value project
```

### 3. Habilitar APIs Necessárias

```bash
make setup-gcloud GCP_PROJECT_ID=seu-projeto-id
```

Ou manualmente:
```bash
gcloud services enable run.googleapis.com
gcloud services enable cloudbuild.googleapis.com
```

### 4. WeatherAPI Key

1. Acesse: https://www.weatherapi.com/signup.aspx
2. Crie uma conta gratuita
3. Copie sua API Key do dashboard
4. Salve em um lugar seguro

---

## Deploy Rápido

### Passo 1: Configurar Variáveis de Ambiente

```bash
# Definir variáveis
export GCP_PROJECT_ID=seu-projeto-gcp
export WEATHER_API_KEY=sua-chave-weatherapi
```

### Passo 2: Deploy do Backend

```bash
make deploy-backend
```

**Saída esperada:**
```
Service [weather-api] revision [weather-api-00001-xxx] has been deployed
and is serving 100 percent of traffic.
Service URL: https://weather-api-xxxxx-uc.a.run.app
```

**Copie a URL do backend!** Você precisará dela no próximo passo.

### Passo 3: Deploy do Frontend

```bash
# Substitua pela URL real do seu backend
export BACKEND_URL=https://weather-api-xxxxx-uc.a.run.app

make deploy-frontend BACKEND_URL=$BACKEND_URL
```

**Saída esperada:**
```
Service [weather-app] revision [weather-app-00001-xxx] has been deployed
and is serving 100 percent of traffic.
Service URL: https://weather-app-xxxxx-uc.a.run.app
```

### Passo 4: Testar

```bash
# Testar backend
curl https://weather-api-xxxxx-uc.a.run.app/weather/01310100

# Abrir frontend no browser
open https://weather-app-xxxxx-uc.a.run.app
```

---

## Deploy Manual Detalhado

### Backend

```bash
cd backend/
export WEATHER_API_KEY=sua-chave-weatherapi

gcloud run deploy weather-api \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080 \
  --set-env-vars="PORT=8080,WEATHER_API_KEY=$WEATHER_API_KEY"

# Copiar URL do serviço
gcloud run services describe weather-api \
  --platform managed \
  --region us-central1 \
  --format 'value(status.url)'
```

---

### Frontend

```bash
cd frontend/

gcloud run deploy weather-app \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080 \
  --set-env-vars="VITE_BACKEND_URL=https://weather-api-xxxxx-uc.a.run.app"
```

---

## Atualizar Deploy Existente

### Backend

```bash
cd backend/
make deploy WEATHER_API_KEY=sua-chave
```

### Frontend

```bash
cd frontend/
make deploy BACKEND_URL=https://weather-api-xxxxx-uc.a.run.app
```

---

## Configurações Avançadas

### Ajustar Recursos (CPU/Memória)

```bash
# Backend
gcloud run services update weather-api \
  --memory 512Mi \
  --cpu 1 \
  --max-instances 10

# Frontend
gcloud run services update weather-app \
  --memory 256Mi \
  --cpu 1 \
  --max-instances 5
```

### Configurar Autoscaling

```bash
gcloud run services update weather-api \
  --min-instances 0 \
  --max-instances 10 \
  --concurrency 80
```

### Adicionar Domínio Customizado

```bash
# Mapear domínio
gcloud run domain-mappings create \
  --service weather-app \
  --domain api.seudominio.com \
  --region us-central1
```

### Ver Logs

```bash
# Backend
gcloud run services logs tail weather-api --region us-central1

# Frontend
gcloud run services logs tail weather-app --region us-central1
```

---

## Troubleshooting

### Erro: "Permission Denied"

```bash
# Configurar permissões
gcloud auth login
gcloud auth configure-docker
```

### Erro: "Service not found"

```bash
# Listar serviços
gcloud run services list --region us-central1
```

### Build falha no Cloud Build

```bash
# Ver logs do build
gcloud builds list --limit 5
gcloud builds log [BUILD_ID]
```

### Frontend não conecta ao Backend

1. Verificar se a URL do backend está correta:
```bash
echo $BACKEND_URL
```

2. Verificar se o backend está respondendo:
```bash
curl $BACKEND_URL/weather/01310100
```

3. Verificar variável de ambiente no Cloud Run:
```bash
gcloud run services describe weather-app \
  --region us-central1 \
  --format 'value(spec.template.spec.containers[0].env)'
```

### CORS Error

O backend já está configurado com CORS. Se ainda houver erro:

1. Verificar se o backend permite a origem do frontend
2. Ver logs do backend para detalhes:
```bash
gcloud run services logs tail weather-api --region us-central1
```

---

## Custos Estimados

**Free Tier do Cloud Run:**
- 2 milhões de requisições/mês
- 360,000 GB-segundos de memória
- 180,000 vCPU-segundos

**Para esta aplicação:**
- Backend: ~$0.50 - $5.00/mês (dependendo do tráfego)
- Frontend: ~$0.30 - $3.00/mês
- WeatherAPI: Grátis até 1 milhão de chamadas/mês

**Dica:** Use `--min-instances 0` para economizar quando não houver tráfego.

---

## Segurança

### Remover Acesso Público (Opcional)

```bash
# Remover --allow-unauthenticated e adicionar IAM
gcloud run services remove-iam-policy-binding weather-api \
  --member="allUsers" \
  --role="roles/run.invoker"
```

### Usar Secret Manager para API Keys

```bash
# Criar secret
echo -n "$WEATHER_API_KEY" | gcloud secrets create weather-api-key --data-file=-

# Usar no Cloud Run
gcloud run services update weather-api \
  --set-secrets="WEATHER_API_KEY=weather-api-key:latest"
```

---

## Monitoramento

### Métricas no Console

1. Acesse: https://console.cloud.google.com/run
2. Selecione seu serviço
3. Ver métricas: Requisições, Latência, Erros

### Alertas

Configure alertas para:
- Alta latência (> 1s)
- Taxa de erro (> 5%)
- Uso de recursos (> 80%)

---

## Limpar Recursos

```bash
# Deletar serviços
gcloud run services delete weather-api --region us-central1 --quiet
gcloud run services delete weather-app --region us-central1 --quiet
```

---

## Referências

- [Cloud Run Documentation](https://cloud.google.com/run/docs)
- [Cloud Build Documentation](https://cloud.google.com/build/docs)
- [Pricing Calculator](https://cloud.google.com/products/calculator)
