# Input Service - Distributed Tracing Demo

Serviço simples que demonstra distributed tracing com OpenTelemetry.

## Propósito

Este serviço existe para demonstrar como trace context é propagado entre microserviços:

```
Cliente → Input-Service → Backend → APIs Externas
          (span raiz)     (child)   (grandchild)
```

Todo o fluxo fica visível no Zipkin como um único trace distribuído.

## Como Funciona

1. Recebe `POST /weather` com CEP no body
2. Inicializa trace com OpenTelemetry + Zipkin exporter
3. Cria HTTP client instrumentado (`otelhttp.NewTransport`)
4. Chama backend propagando trace context via HTTP headers
5. Retorna resposta do backend ao cliente

## Tecnologias

- Go 1.23
- OpenTelemetry (OTEL)
- Zipkin exporter
- HTTP client instrumentation

## Configuração

**Environment variables:**
```env
PORT=8000
OTEL_SERVICE_NAME=input-service
ZIPKIN_ENDPOINT=http://zipkin:9411/api/v2/spans
BACKEND_URL=http://backend:3000/weather
```

## API

### POST /weather

**Request Body:**
```json
{
  "cep": "01310100"
}
```

**Response:** Mesma estrutura do backend

**Exemplo:**
```bash
curl -X POST http://localhost:8000/weather \
  -H "Content-Type: application/json" \
  -d '{"cep":"01310100"}'
```

## Testar Distributed Tracing

```bash
# 1. Enviar request
make test-distributed-tracing

# 2. Visualizar no Zipkin
make zipkin
# Ou: http://localhost:9411

# 3. Clicar em "Run Query" e ver o trace completo
```

Você verá:
- **input-service** (span raiz)
  - **backend/http-server** (child span)
    - **ViaCEP.GetAddress** (grandchild)
    - **WeatherAPI.GetTemperature** (grandchild)

## Código Principal

O serviço inteiro está em um único arquivo: `main.go`

**Instrumentação key:**
- `initTelemetry()`: Configura OTEL com Zipkin
- `otelhttp.NewTransport()`: Instrumenta HTTP client para propagação automática
- `http.NewRequestWithContext()`: Passa context com trace info

Não é necessário passar headers manualmente - o `otelhttp` faz isso automaticamente.
