# Weather API - Backend

API para consulta de temperatura por CEP (Código Postal Brasileiro).

## 🏗️ Arquitetura

- **Clean Architecture** com separação de camadas
- **Dependency Injection** com Google Wire
- **HTTP Router** com chi
- **Configuration** com Viper

## 📋 Pré-requisitos

- Go 1.23+
- Docker (para containerização)
- Google Cloud SDK (para deploy)
- Wire CLI: `go install github.com/google/wire/cmd/wire@latest`

## 🚀 Como usar

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

### Deploy no Cloud Run

```bash
# Configurar variáveis (substitua com seu projeto)
export GCP_PROJECT_ID=seu-projeto-gcp
export WEATHER_API_KEY=sua-chave

# Deploy completo (build + push + deploy)
make deploy-dev
```

## 📡 Endpoints

### GET /weather/{zipcode}

Retorna a temperatura atual para o CEP informado.

**Exemplo:**
```bash
curl http://localhost:3000/weather/01310100
```

**Resposta de sucesso (200):**
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

**Erros:**
- `422` - CEP inválido: `{"message": "invalid zipcode"}`
- `404` - CEP não encontrado: `{"message": "can not find zipcode"}`
- `500` - Erro interno: `{"message": "internal server error"}`

## 🧪 Testes

```bash
# Rodar todos os testes
make test

# Com coverage
make test-coverage
```

## 📦 Estrutura do Projeto

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

## 🔧 Comandos Makefile

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
| `make deploy` | Deploy no Cloud Run |
| `make dev` | Wire + Build + Run |

## 📝 Tecnologias

- **Go 1.23**
- **chi** - HTTP Router
- **Viper** - Configuration
- **Wire** - Dependency Injection
- **WeatherAPI** - Dados de temperatura
- **ViaCEP** - Consulta de CEP
