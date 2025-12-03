# Rate Limiter API

Sistema de rate limiting configurável em Go com suporte a limitação por IP e token, usando Redis para armazenamento distribuído.

## 🚀 Features

- ✅ **Rate limiting por IP**: Limita requisições por endereço IP
- ✅ **Rate limiting por Token**: Suporta diferentes limites por tier (basic/premium)
- ✅ **Prioridade de Token**: Limites de token sobrescrevem limites de IP
- ✅ **Janela Configurável**: Window de tempo personalizável (padrão: 1 segundo)
- ✅ **Bloqueio Temporário**: Bloqueia IPs/tokens que excedem o limite
- ✅ **Redis Backend**: Armazenamento distribuído com operações atômicas
- ✅ **Chi Router**: Framework HTTP idiomático e performático
- ✅ **Docker Ready**: Deploy simplificado com Docker Compose

---

## 📋 Índice

- [Configuração](#-configuração)
- [Como Usar](#-como-usar)
- [API Endpoints](#-api-endpoints)
- [Testes](#-testes)
- [Troubleshooting](#-troubleshooting)

---

### Componentes

- **Middleware**: Intercepta requisições HTTP, extrai IP e token
- **Limiter**: Implementa lógica de rate limiting
- **Storage**: Interface para persistência (Redis)
- **Config**: Carrega configurações

---

## ⚙️ Configuração

### Variáveis de Ambiente

| Variável | Descrição | Padrão | Exemplo |
|----------|-----------|--------|---------|
| `REDIS_ADDR` | Endereço do Redis | - | `redis:6379` |
| `RATE_LIMIT_IP` | Limite de requisições por IP (por window) | - | `10` |
| `RATE_LIMIT_TOKEN_BASIC` | Limite para tokens basic | - | `50` |
| `RATE_LIMIT_TOKEN_PREMIUM` | Limite para tokens premium | - | `200` |
| `RATE_LIMIT_WINDOW` | Janela de tempo em segundos | `1` | `1` |
| `RATE_LIMIT_IP_BLOCK_DURATION` | Tempo de bloqueio IP (segundos) | - | `300` |
| `RATE_LIMIT_TOKEN_BLOCK_DURATION` | Tempo de bloqueio token (segundos) | - | `300` |
| `SERVER_PORT` | Porta do servidor | - | `8080` |

### Arquivo .env

Copie `.env.example` para `.env` e ajuste conforme necessário:

```bash
cp .env.example .env
```

---

## 🚀 Como Usar

### Opção 1: Docker Compose (Recomendado)

```bash
# Iniciar todos os serviços
make up

# Verificar logs
make logs

# Parar serviços
make down
```

### Opção 2: Desenvolvimento Local

```bash
# Iniciar apenas Redis
make dev-up

# Rodar aplicação localmente
make run-local

# Ou diretamente:
go run cmd/server/main.go
```

### Comandos Úteis (Makefile)

```bash
make help              # Lista todos os comandos
make build            # Build das imagens Docker
make restart          # Reinicia serviços
make redis-cli        # Acessa Redis CLI
make redis-keys       # Lista todas as keys de rate limit
make test-ip          # Testa rate limit por IP
make test-token TOKEN=xxx  # Testa rate limit por token
make health           # Verifica saúde dos serviços
```

---

## 🔌 API Endpoints

### GET /

Endpoint de teste para verificar rate limiting.

**Response:**
```
200 OK - Requisição permitida
429 Too Many Requests - Limite excedido
```

**Exemplo:**
```bash
curl http://localhost:8080/
```

---

### GET /auth?tier={basic|premium}

Gera um token de acesso para teste.

**Query Parameters:**
- `tier` (opcional): `basic` ou `premium` (padrão: `basic`)

**Response:**
```json
{
  "token": "premium:550e8400-e29b-41d4-a716-446655440000",
  "tier": "premium"
}
```

**Exemplo:**
```bash
# Token básico
curl "http://localhost:8080/auth?tier=basic"

# Token premium
curl "http://localhost:8080/auth?tier=premium"
```

---

### Usando Tokens

Inclua o token no header `API_KEY`:

```bash
TOKEN=$(curl -s "http://localhost:8080/auth?tier=premium" | jq -r .token)
curl -H "API_KEY: $TOKEN" http://localhost:8080/
```

---

## 🧪 Testes

### Teste 1: Rate Limit por IP

```bash
# Envia 15 requisições (limite é 10)
make test-ip
```

**Resultado Esperado:**
- Primeiras 10 requisições: `200 OK`
- Próximas 5 requisições: `429 Too Many Requests`

---

### Teste 2: Rate Limit por Token (Basic)

```bash
# Obter token
TOKEN=$(curl -s "http://localhost:8080/auth?tier=basic" | jq -r .token)

# Testar com 55 requisições (limite é 50)
make test-token TOKEN=$TOKEN
```

**Resultado Esperado:**
- Primeiras 50 requisições: `200 OK`
- Próximas 5 requisições: `429 Too Many Requests`

---

### Teste 3: Rate Limit por Token (Premium)

```bash
TOKEN=$(curl -s "http://localhost:8080/auth?tier=premium" | jq -r .token)

# Enviar muitas requisições
for i in {1..210}; do
  curl -s -o /dev/null -w "Request $i: %{http_code}\n" \
    -H "API_KEY: $TOKEN" \
    http://localhost:8080/
done
```

**Resultado Esperado:**
- Primeiras 200 requisições: `200 OK`
- Próximas 10 requisições: `429 Too Many Requests`

---

### Teste 4: Prioridade Token > IP

```bash
# Sem token (limite IP = 10)
for i in {1..12}; do curl http://localhost:8080/; done

# Com token premium (limite = 200, ignora IP)
TOKEN=$(curl -s "http://localhost:8080/auth?tier=premium" | jq -r .token)
for i in {1..15}; do curl -H "API_KEY: $TOKEN" http://localhost:8080/; done
```

**Resultado:** Mesmo IP bloqueado, token premium permite 200 requisições.

---

### Teste 5: Requisições Concorrentes

```bash
make test-concurrent
```

**Resultado:** Redis INCR atômico garante que exatamente 10 requisições passam, mesmo com concorrência.

---

### Teste 6: Janela de Tempo Customizada

Edite `.env`:
```env
RATE_LIMIT_WINDOW=5  # 5 segundos
RATE_LIMIT_IP=10     # 10 requisições em 5 segundos = 2 req/s
```

```bash
make restart

# Enviar 10 requisições em < 5 segundos
for i in {1..10}; do curl http://localhost:8080/; done

# Aguardar 6 segundos
sleep 6

# Enviar mais 10 (deve funcionar, window resetou)
for i in {1..10}; do curl http://localhost:8080/; done
```

---

## 🔍 Debugging

### Ver Keys no Redis

```bash
make redis-keys

# Ou manualmente:
make redis-cli
> KEYS ratelimit:*
```

**Formato das Keys:**
```
ratelimit:counter:ip:192.168.1.1        # Contador de requisições
ratelimit:block:ip:192.168.1.1          # Flag de bloqueio
ratelimit:counter:token:premium:abc-123  # Contador de token
ratelimit:block:token:premium:abc-123    # Bloqueio de token
```

---

### Ver TTL de uma Key

```bash
make redis-cli
> TTL ratelimit:counter:ip:192.168.1.1
```

**Retorno:**
- Número positivo: segundos restantes até expirar
- `-1`: key existe mas sem expiração
- `-2`: key não existe

---

### Limpar Redis (Reset)

```bash
make redis-flush
```

⚠️ **CUIDADO:** Isso apaga TODOS os dados do Redis!

---

## 🐛 Troubleshooting

### Problema: Todas as requisições retornam 429

**Causa:** IP pode estar bloqueado de execução anterior.

**Solução:**
```bash
make redis-flush
make restart
```

---

### Problema: Config não está sendo aplicada

**Causa:** Precedência de variáveis de ambiente.

**Solução:**
1. Verifique `docker-compose.yml` seção `environment:`
2. Essas variáveis sobrescrevem `.env`
3. Use `docker-compose config` para ver config final:
   ```bash
   docker-compose config
   ```

---

### Problema: "connection refused" ao Redis

**Causa:** Redis não está rodando ou app não consegue conectar.

**Solução:**
```bash
# Verificar status
make status

# Verificar logs
make logs-redis

# Para desenvolvimento local, use:
REDIS_ADDR=localhost:6379 go run cmd/server/main.go
```

---

### Problema: App não inicia no Docker

**Causa:** Erro de build ou configuração.

**Solução:**
```bash
# Ver logs detalhados
docker-compose logs app

# Rebuild completo
make clean
make build
make up
```

---

## 📊 Monitoramento

### Verificar Saúde dos Serviços

```bash
make health
```

---

### Logs em Tempo Real

```bash
# App
make logs

# Redis
make logs-redis

# Ambos
make watch
```

---

### Estatísticas do Redis

```bash
make redis-cli
> INFO stats
> INFO keyspace
```

---

## 🏗️ Estrutura do Projeto

```
.
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── config/
│   └── config.go                # Configuração com Viper
├── internal/
│   ├── limiter/
│   │   └── limiter.go           # Lógica de rate limiting
│   ├── middleware/
│   │   └── ratelimit.go         # HTTP middleware
│   └── storage/
│       ├── storage.go           # Interface
│       └── redis.go             # Implementação Redis
├── .env                         # Variáveis de ambiente
├── docker-compose.yml           # Orquestração Docker
├── Dockerfile                   # Build da aplicação
├── Makefile                     # Comandos úteis
├── go.mod                       # Dependências
└── README.md                    # Este arquivo
```

---

## 🔒 Comportamento de Rate Limiting

### Algoritmo

1. **Extrair identificadores** (IP + opcional token do header `API_KEY`)
2. **Determinar limite:**
   - Se token presente → usar limite do tier do token
   - Senão → usar limite de IP
3. **Verificar se bloqueado** (de violação anterior)
   - Se bloqueado → retornar 429
4. **Incrementar contador** (Redis INCR atômico)
   - Contador tem TTL = `RATE_LIMIT_WINDOW`
5. **Verificar se excedeu:**
   - Se `count > limit` → bloquear por `BLOCK_DURATION`
   - Senão → permitir requisição

---

### Sliding Window

```
RATE_LIMIT_WINDOW=1 segundo
RATE_LIMIT_IP=10

Requisições:
T=0.0s → count=1 ✅
T=0.1s → count=2 ✅
...
T=0.9s → count=10 ✅
T=1.0s → count=11 ❌ (BLOCKED for 300s)

T=1.5s → count=1 (novo window) ✅
```

---

### Formato de Tokens

Tokens seguem o padrão: `{tier}:{uuid}`

**Exemplos:**
```
basic:550e8400-e29b-41d4-a716-446655440000
premium:6ba7b810-9dad-11d1-80b4-00c04fd430c8
```

Parsing:
- Parte antes do `:` = tier
- Parte depois do `:` = identificador único
- Se tier inválido → default para `basic`

---

## 📝 Licença

MIT

---

## 👤 Autor

Desenvolvido como parte do desafio Full Cycle Go Expert.

---

## ⭐ Links Úteis

- [Go Chi Router](https://github.com/go-chi/chi)
- [Redis Documentation](https://redis.io/docs/)
- [Viper Configuration](https://github.com/spf13/viper)
- [Docker Compose](https://docs.docker.com/compose/)
