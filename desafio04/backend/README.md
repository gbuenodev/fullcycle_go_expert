# Weather API - Backend

API para consulta de temperatura por CEP (Código Postal Brasileiro).

## Arquitetura

- **Clean Architecture** com separação de camadas
- **Dependency Injection** com Google Wire
- **HTTP Router** com chi
- **Configuration** com Viper

## Pré-requisitos

- Go 1.23+
- Docker (para containerização)
- Wire CLI: `go install github.com/google/wire/cmd/wire@latest`

## Como usar

### Configuração

1. Copie o arquivo de exemplo:
```bash
cp .env.example .env
```

2. Edite o `.env` com suas credenciais:
```env
PORT=3000
WEATHER_API_KEY=sua-chave-aqui
WEATHER_API_BASE_URL=https://api.weatherapi.com/v1
VIACEP_BASE_URL=https://viacep.com.br/ws
```

### Desenvolvimento Local

```bash
# Ver todos os comandos disponíveis
make help

# Gerar código Wire + Build + Run
make dev

# Apenas rodar (sem rebuild)
make run

# Rodar testes
make test

# Rodar testes com coverage
make test-coverage
```

### Docker

```bash
# Build da imagem
make docker-build

# Rodar container localmente
make docker-run
```

## Endpoints

### POST /weather

Retorna a temperatura atual para o CEP informado.

**Request Body:**
```json
{
  "cep": "01310100"
}
```

**Exemplo:**
```bash
curl -X POST http://localhost:3000/weather \
  -H "Content-Type: application/json" \
  -d '{"cep":"01310100"}'
```

**Resposta de sucesso (200):**
```json
{
  "city": "São Paulo",
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

**Erros:**
- `400` - Request body inválido: `{"message": "invalid request body"}`
- `422` - CEP inválido: `{"message": "invalid zipcode"}`
- `404` - CEP não encontrado: `{"message": "can not find zipcode"}`
- `500` - Erro interno: `{"message": "internal server error"}`

## Observabilidade

### OpenTelemetry & Distributed Tracing

O backend possui instrumentação completa com OpenTelemetry:

**Configuração via environment variables:**
```env
OTEL_SERVICE_NAME=weather-api
ZIPKIN_ENDPOINT=http://zipkin:9411/api/v2/spans
```

**Spans criados automaticamente:**
- HTTP requests (via middleware `otelhttp`)
- Chamadas ao ViaCEP
- Chamadas ao WeatherAPI

**Trace context propagation:**
O backend recebe e propaga trace context via HTTP headers, permitindo distributed tracing quando chamado pelo input-service.

### Métricas (Prometheus)

**Endpoint:** `GET /metrics`

Expõe métricas no formato Prometheus:
```bash
curl http://localhost:3000/metrics
```

Prometheus configurado para scraping automático (veja `prometheus.yml` na raiz do projeto).

## Testes

```bash
# Rodar todos os testes
make test

# Com coverage
make test-coverage
```

## Estrutura do Projeto

```
backend/
├── cmd/
│   └── server/          # Entry point
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── config/              # Configurações
├── internal/
│   ├── dto/            # Data Transfer Objects
│   ├── entity/         # Entidades de domínio
│   ├── gateway/        # Interfaces (portas)
│   ├── usecase/        # Lógica de negócio
│   └── infra/
│       ├── api/        # HTTP Handlers
│       ├── clients/    # HTTP Clients externos
│       └── webserver/  # Router e WebServer
└── pkg/
    └── validator/      # Validações reutilizáveis
```

## Comandos Makefile

| Comando | Descrição |
|---------|-----------|
| `make help` | Mostra todos os comandos |
| `make build` | Compila o binário |
| `make run` | Executa a aplicação |
| `make test` | Roda os testes |
| `make wire` | Gera código Wire |
| `make clean` | Remove artifacts |
| `make docker-build` | Build da imagem Docker |
| `make docker-run` | Roda container localmente |
| `make dev` | Wire + Build + Run |

## Tecnologias

- **Go 1.23**
- **chi** - HTTP Router
- **Viper** - Configuration
- **Wire** - Dependency Injection
- **OpenTelemetry** - Observability
- **Zipkin** - Distributed Tracing
- **Prometheus** - Metrics
- **WeatherAPI** - Dados de temperatura
- **ViaCEP** - Consulta de CEP
