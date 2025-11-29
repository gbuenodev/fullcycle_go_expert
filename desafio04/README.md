# Weather App - Consulta de Clima por CEP

Aplicação fullstack para consultar informações de clima baseado em CEP brasileiro com **observabilidade distribuída completa**.

## Features

- Consulta de clima por CEP brasileiro
- Integração com ViaCEP + WeatherAPI
- Distributed Tracing com Zipkin
- Métricas com Prometheus
- OpenTelemetry Instrumentation
- Múltiplos microserviços
- Clean Architecture no backend
- UI moderna com Material-UI

DEMOS DISPONÍVEIS EM:
- Front: [APP](https://weather-app-watebi5u2q-uc.a.run.app/)
- Back: [API](https://weather-api-watebi5u2q-uc.a.run.app/)

** Veja rotas disponíveis na API [aqui](#API_Endpoints).

## Arquitetura

```
┌─────────────────┐
│   Frontend      │  React + TypeScript + MUI
│   Port: 8080    │
└────────┬────────┘
         │
         │
         ▼
      ┌─────────────────┐
      │   Backend       │  Go + Clean Architecture
      │   Port: 3000    │
      └────┬────────┬───┘
           │        │
           ▼        ▼
        ViaCEP  WeatherAPI


Distributed Tracing Demo:

┌─────────────────┐
│ Input Service   │
│   Port: 8000    │
└────────┬────────┘
         │
         ▼
      ┌─────────────────┐
      │   Backend       │
      │   Port: 3000    │
      └────┬────────┬───┘
           │        │
           ▼        ▼
        ViaCEP  WeatherAPI

────────────────────────────────────────
     Observabilidade
────────────────────────────────────────

┌──────────────┐        ┌──────────────┐
│   Zipkin     │        │  Prometheus  │
│  Port: 9411  │        │  Port: 9090  │
└──────────────┘        └──────────────┘
```

## Stack

### Backend (weather-api)
- **Go 1.23**
- **Clean Architecture** (Entity, UseCase, Gateway, Infra)
- **Dependency Injection** com Google Wire
- **Router:** Chi
- **Config:** Viper
- **OpenTelemetry** com Zipkin exporter
- **APIs Externas:** ViaCEP + WeatherAPI

### Input Service (Demo)
- **Go 1.23**
- **OpenTelemetry** instrumentation
- **HTTP Client** instrumentado para distributed tracing
- Propagação automática de trace context
- Usado apenas para demonstrar distributed tracing

### Frontend
- **React 18**
- **TypeScript**
- **Material-UI (MUI)**
- **Vite**
- **Axios**
- **Zod** para validação

### Observabilidade
- **Zipkin** - Distributed Tracing
- **Prometheus** - Metrics & Monitoring
- **OpenTelemetry** - Instrumentation padrão

## Quick Start

### Pré-requisitos

- Docker e Docker Compose
- Make (opcional, mas recomendado)

### Subir ambiente completo

```bash
# Ver todos os comandos disponíveis
make help

# Build e subir tudo
make quickstart

# Ou manualmente
docker-compose up -d
```

### Acessar serviços

**Aplicações:**
- Frontend: http://localhost:8080
- Backend: http://localhost:3000
- Input Service: http://localhost:8000

**Observabilidade:**
- Zipkin (Tracing): http://localhost:9411
- Prometheus (Metrics): http://localhost:9090
- Backend Metrics: http://localhost:3000/metrics

### Testar Distributed Tracing

```bash
# Enviar request que flui por todos os serviços
make test-distributed-tracing

# Depois acesse Zipkin e veja o trace completo
make zipkin
```

No Zipkin você verá:
```
input-service (span raiz)
  └─ HTTP POST → backend
      └─ weather-api: http-server
          ├─ ViaCEP.GetAddress
          └─ WeatherAPI.GetTemperature
```

## Comandos Make

### Gerenciamento de Serviços

```bash
make up                    # Sobe todos os serviços
make down                  # Para e remove containers
make restart               # Reinicia serviços
make rebuild               # Rebuild completo
make ps                    # Status dos containers
make health                # Verifica health de todos
```

### Logs

```bash
make logs                  # Todos os logs
make logs-backend          # Logs do backend
make logs-frontend         # Logs do frontend
make logs-input            # Logs do input-service
make logs-zipkin           # Logs do Zipkin
make logs-prometheus       # Logs do Prometheus
```

### Testes

```bash
make test-backend-post     # Testa POST /weather do backend
make test-input-service    # Testa input-service
make test-distributed-tracing  # Testa tracing distribuído
make test-all              # Executa todos os testes
```

### Observabilidade

```bash
make zipkin                # Abre Zipkin UI
make prometheus            # Abre Prometheus UI
make metrics               # Mostra métricas do backend
```

### Desenvolvimento

```bash
make dev-backend           # Shell no container backend
make dev-frontend          # Shell no container frontend
make dev-input             # Shell no container input-service
```

### Limpeza

```bash
make clean                 # Remove containers e volumes
make clean-all             # Remove tudo (imagens também)
make prune                 # Limpa sistema Docker
```

## API Endpoints

### Backend (Port 3000)

#### POST /weather
Consulta clima por CEP via JSON body.

**Request:**
```bash
curl -X POST http://localhost:3000/weather \
  -H "Content-Type: application/json" \
  -d '{"cep":"01310100"}'
```

**Responses:**
- `200 OK` - Sucesso
- `400 Bad Request` - JSON inválido
- `404 Not Found` - CEP não encontrado
- `422 Unprocessable Entity` - CEP inválido

### Input Service (Port 8000) - Distributed Tracing Demo

#### POST /weather
Demo de distributed tracing. Recebe CEP, chama backend, propaga trace context.

**Request:**
```bash
curl -X POST http://localhost:8000/weather \
  -H "Content-Type: application/json" \
  -d '{"cep":"01310100"}'
```

**Response:** Mesma estrutura do backend.

## Observabilidade

### Distributed Tracing (Zipkin)

**Como funciona:**

1. Request chega no **input-service**
2. Input-service **cria span raiz** e trace ID
3. Chama backend **propagando trace context** via HTTP headers
4. Backend **continua o mesmo trace**
5. Backend cria **child spans** para ViaCEP e WeatherAPI
6. **Todos os spans são enviados para Zipkin**

**Visualizar traces:**

```bash
# 1. Enviar request
make test-distributed-tracing

# 2. Abrir Zipkin
make zipkin

# 3. Clicar em "Run Query"
# 4. Clicar no trace mais recente
```

**O que você verá no Zipkin:**

- Timeline completa do request
- Tempo de cada operação
- Hierarquia de spans (parent-child)
- Tags e atributos (CEP, cidade, temperatura)
- Erros (se houver)

### Métricas (Prometheus)

**Endpoint de métricas:**

```bash
# Ver métricas do backend
curl http://localhost:3000/metrics

# Ou via Make
make metrics
```

**Prometheus UI:**

```bash
make prometheus
# ou
open http://localhost:9090
```

**Queries úteis:**

```promql
# Targets sendo monitorados
up

# Métricas HTTP (se configurado)
http_request_duration_seconds
```

## Configuração OpenTelemetry

### Backend

**Inicialização automática** em `main.go`:
```go
shutdown, err := observability.InitTelemetry(ctx)
defer shutdown(ctx)
```

**Trace provider:**
- Exporter: Zipkin HTTP
- Sampling: 100% (AlwaysSample)
- Resource: service.name="weather-api"

**Spans criados:**
- HTTP requests (automático via `otelhttp`)
- ViaCEP calls (manual)
- WeatherAPI calls (manual)

### Input Service

**HTTP Client instrumentado:**
```go
client := &http.Client{
    Transport: otelhttp.NewTransport(http.DefaultTransport),
}
```

**Propagação automática** de trace context via HTTP headers:
- `traceparent`
- `tracestate`

## Testes

### Health Checks

```bash
make health
```

Verifica:
- Backend (port 3000)
- Frontend (port 8080)
- Zipkin (port 9411)
- Prometheus (port 9090)

### Teste Manual - Distributed Tracing

```bash
# 1. Request no input-service
curl -X POST http://localhost:8000/weather \
  -H "Content-Type: application/json" \
  -d '{"cep":"01310100"}'

# 2. Ver trace no Zipkin
open http://localhost:9411
```

### Teste Manual - Backend Direto

```bash
# GET
curl http://localhost:3000/weather/01310100

# POST
curl -X POST http://localhost:3000/weather \
  -H "Content-Type: application/json" \
  -d '{"cep":"01310100"}'
```

## Desenvolvimento

### Estrutura de Diretórios

```
desafio04/
├── backend/                # Go backend (Clean Architecture)
│   ├── cmd/server/        # Main + Wire
│   ├── config/            # Configuração (Viper)
│   ├── internal/
│   │   ├── dto/           # Data Transfer Objects
│   │   ├── entity/        # Domain entities
│   │   ├── gateway/       # Interfaces
│   │   ├── infra/
│   │   │   ├── api/       # HTTP handlers
│   │   │   ├── clients/   # External API clients
│   │   │   └── webserver/ # Router setup
│   │   ├── observability/ # OTEL setup
│   │   └── usecase/       # Business logic
│   └── pkg/               # Shared utilities
│
├── input-service/         # Simple Go service
│   ├── main.go           # Single file service
│   ├── Dockerfile
│   └── go.mod
│
├── frontend/             # React frontend
│   ├── src/
│   │   ├── components/   # React components
│   │   ├── context/      # Context providers
│   │   ├── services/     # API calls
│   │   └── theme/        # MUI theme
│   └── package.json
│
├── docker-compose.yml    # Orquestração de serviços
├── prometheus.yml        # Config do Prometheus
├── Makefile             # Comandos úteis
└── README.md            # Este arquivo
```

### Adicionar Nova Instrumentação

**No backend:**

```go
import "go.opentelemetry.io/otel"

func MinhaFuncao(ctx context.Context) {
    tracer := otel.Tracer("meu-componente")
    ctx, span := tracer.Start(ctx, "MinhaFuncao")
    defer span.End()

    span.SetAttributes(attribute.String("meu-atributo", "valor"))

    // ... seu código ...

    if err != nil {
        span.RecordError(err)
        span.SetStatus(codes.Error, "descrição do erro")
    }
}
```

## Troubleshooting

### Serviços não sobem

```bash
# Ver logs
make logs

# Rebuild sem cache
make build-no-cache
make up
```

### Traces não aparecem no Zipkin

```bash
# Verificar se Zipkin está rodando
docker-compose ps zipkin

# Ver logs do Zipkin
make logs-zipkin

# Verificar env vars do backend
docker-compose exec backend env | grep ZIPKIN
```

### Prometheus não mostra métricas

```bash
# Verificar targets
open http://localhost:9090/targets

# Ver logs
make logs-prometheus

# Testar endpoint de métricas
curl http://localhost:3000/metrics
```

### Porta já em uso

```bash
# Ver o que está usando as portas
lsof -i :3000  # Backend
lsof -i :8000  # Input-service
lsof -i :8080  # Frontend
lsof -i :9411  # Zipkin
lsof -i :9090  # Prometheus

# Parar tudo
make down
```

## Diferenciais

- Distributed Tracing production-ready
- OpenTelemetry (padrão da indústria)
- Propagação automática de trace context
- Visibilidade completa do fluxo de requests
- Instrumentação em múltiplas camadas (HTTP, use case, clients)
- Observabilidade de microserviços
- Clean Architecture no backend
- Makefile organizado com todos os comandos
- Docker Compose para ambiente completo

## Documentação Adicional

- [Backend README](backend/README.md) - Detalhes do backend
- [Frontend README](frontend/README.md) - Detalhes do frontend
- [OpenTelemetry Docs](https://opentelemetry.io/docs/languages/go/)
- [Zipkin Docs](https://zipkin.io/)
- [Prometheus Docs](https://prometheus.io/docs/introduction/overview/)

## Licença

MIT
