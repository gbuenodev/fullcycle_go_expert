# Desafio 05 - Sistema de Leilões (Auction System)

API RESTful para gerenciamento de leilões online com Clean Architecture e MongoDB.

## 📋 Sobre o Projeto

Sistema de leilões que permite criar e gerenciar leilões, fazer lances e consultar informações de usuários. Desenvolvido seguindo os princípios de Clean Architecture para garantir separação de responsabilidades e manutenibilidade.

## 🚀 Features

- **Gerenciamento de Leilões**
  - Criar novos leilões
  - Listar leilões (com filtros opcionais)
  - Buscar leilão por ID
  - Consultar vencedor de um leilão

- **Sistema de Lances**
  - Criar lances em leilões ativos
  - Listar lances de um leilão
  - Validação de status do leilão (impede lances em leilões encerrados)

- **Usuários**
  - Criar novo usuário
  - Consultar informações de usuário

- **Validações de Negócio**
  - **Auction:**
    - ProductName não pode ser vazio
    - Category mínimo 3 caracteres
    - Description entre 10-50 caracteres
    - Condition e Status validados (enum)
  - **Bid:**
    - Leilão deve existir
    - Leilão deve estar ativo (status = 0)
    - Retorna 403 Forbidden se tentar lance em leilão encerrado
  - **User:**
    - Name não pode ser vazio
    - Name mínimo 2 caracteres

## 🏗️ Arquitetura

```
desafio05/
├── cmd/auction/              # Entry point
│   ├── main.go              # Inicialização e DI
│   └── .env                 # Configurações
├── configs/                  # Configurações e utilidades
│   ├── database/mongodb/    # Conexão MongoDB
│   ├── logger/              # Logger (Zap)
│   └── rest_err.go/         # Error handling
├── internal/
│   ├── entity/              # Entidades de domínio
│   │   ├── auction_entity/  # Auction, validações
│   │   ├── bid_entity/      # Bid
│   │   └── user_entity/     # User
│   ├── usecase/             # Regras de negócio
│   │   ├── auction_usecase/
│   │   ├── bid_usecase/
│   │   └── user_usecase/
│   ├── infra/
│   │   ├── repository/      # Implementação de persistência
│   │   │   ├── auction/
│   │   │   ├── bid/
│   │   │   └── user/
│   │   └── api/web/         # Camada HTTP
│   │       ├── controller/  # Controllers (Gin)
│   │       └── validation/  # Validação de requests
│   └── internal_errors/     # Errors customizados
├── docker-compose.yml       # Orquestração
├── Dockerfile              # Build da aplicação
├── Makefile               # Comandos úteis
└── README.md             # Este arquivo
```

## 🛠️ Stack Tecnológica

- **Go 1.25+**
- **MongoDB** - Banco de dados NoSQL
- **Gin** - Web framework
- **Zap** - Structured logging
- **godotenv** - Gerenciamento de variáveis de ambiente
- **Docker & Docker Compose** - Containerização

## 📦 Pré-requisitos

- **Docker** e **Docker Compose** instalados
- **Make** (opcional, mas recomendado)
- **Go 1.25+** (apenas para desenvolvimento local)

## 🚀 Quick Start

### Opção 1: Usando Docker (Recomendado)

```bash
# Ver todos os comandos disponíveis
make help

# Subir ambiente completo (MongoDB + API containerizados)
make run

# Ver logs da API
make docker-logs
```

A API estará disponível em: `http://localhost:8080`

### Opção 2: Desenvolvimento Local

```bash
# Subir MongoDB e rodar app localmente (hot reload)
make dev
```

## 📡 API Endpoints

### Auctions

#### POST /auctions
Criar um novo leilão.

**Request Body:**
```json
{
  "productName": "iPhone 15 Pro",
  "category": "Eletrônicos",
  "description": "iPhone 15 Pro 256GB Azul Titânio, estado de novo",
  "condition": 0
}
```

**Validações:**
- `productName`: não pode ser vazio
- `category`: mínimo 3 caracteres
- `description`: entre 10 e 50 caracteres
- `condition`: 0 (New), 1 (Used), 2 (Refurbished)

**Response:** `201 Created`

---

#### GET /auctions
Listar leilões (com filtros opcionais).

**Query Parameters (opcionais):**
- `status`: 0 (Active) ou 1 (Completed)
- `category`: filtrar por categoria
- `productName`: filtrar por nome do produto

**Exemplo:**
```bash
curl "http://localhost:8080/auctions?status=0&category=Eletrônicos"
```

**Response:** `200 OK`
```json
[
  {
    "id": "...",
    "productName": "iPhone 15 Pro",
    "category": "Eletrônicos",
    "description": "iPhone 15 Pro 256GB...",
    "condition": 0,
    "status": 0,
    "timestamp": "2024-12-01T10:00:00Z"
  }
]
```

---

#### GET /auctions/:auctionId
Buscar leilão por ID.

**Response:** `200 OK`

---

#### GET /auctions/winner/:auctionId
Consultar vencedor de um leilão.

**Response:** `200 OK`

---

### Bids

#### POST /bid
Criar um lance em um leilão.

**Request Body:**
```json
{
  "userId": "user-id-here",
  "auctionId": "auction-id-here",
  "amount": 1500.00
}
```

**Possíveis Respostas:**
- `201 Created` - Lance criado com sucesso
- `403 Forbidden` - Leilão encerrado
  ```json
  {
    "message": "auction is closed",
    "err": "forbidden",
    "code": 403
  }
  ```
- `400 Bad Request` - Dados inválidos
- `404 Not Found` - Leilão não encontrado

---

#### GET /bid/:auctionId
Listar todos os lances de um leilão.

**Response:** `200 OK`

---

### Users

#### POST /user
Criar um novo usuário.

**Request Body:**
```json
{
  "name": "João Silva"
}
```

**Validações:**
- `name`: não pode ser vazio
- `name`: mínimo 2 caracteres

**Response:** `201 Created`
```json
{
  "id": "generated-user-id",
  "name": "João Silva"
}
```

---

#### GET /user/:userId
Buscar informações de um usuário.

**Response:** `200 OK`
```json
{
  "id": "user-id",
  "name": "João Silva"
}
```

---

### Health Check

#### GET /health
Verificar status da API.

**Response:** `200 OK`
```json
{
  "status": "ok"
}
```

---

## 🔧 Comandos Make

### Principais
```bash
make run               # Subir tudo com Docker (MongoDB + API containerizados)
make dev               # Desenvolvimento local (MongoDB no Docker, app local)
make build             # Compilar binário Go
make test              # Executar todos os testes
make clean             # Remover tudo (containers, volumes, binários)
```

### Docker (auxiliares)
```bash
make docker-up         # Subir serviços
make docker-down       # Parar serviços
make docker-build      # Rebuild de imagens sem cache
make docker-logs       # Ver logs da API
make docker-clean      # Remover volumes
```

**Comandos Go básicos** (use diretamente):
```bash
go mod tidy            # Organizar dependências
go test -v ./...       # Executar testes
go run cmd/auction/main.go  # Rodar app
```

## 🧪 Testando a API

### Health Check
```bash
curl http://localhost:8080/health
```

### Criar um usuário
```bash
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name": "João Silva"}'
```

### Criar um leilão
```bash
curl -X POST http://localhost:8080/auctions \
  -H "Content-Type: application/json" \
  -d '{
    "productName": "MacBook Pro M3",
    "category": "Eletrônicos",
    "description": "MacBook Pro 14 M3 Pro 18GB 512GB Space Black",
    "condition": 0
  }'
```

### Listar leilões ativos
```bash
curl "http://localhost:8080/auctions?status=0"
```

### Criar um lance
```bash
curl -X POST http://localhost:8080/bid \
  -H "Content-Type: application/json" \
  -d '{
    "userId": "user-id-aqui",
    "auctionId": "auction-id-aqui",
    "amount": 5000.00
  }'
```

### Consultar vencedor de um leilão
```bash
curl http://localhost:8080/auctions/winner/auction-id-aqui
```

## 🗄️ Variáveis de Ambiente

Configuradas em `cmd/auction/.env`:

```env
MONGODB_URL=mongodb://localhost:27017
MONGODB_DB=auctions
AUCTION_DURATION=24h
BATCH_INSERT_INTERVAL=7m
MAX_BATCH_SIZE=10
```

**Descrição das variáveis:**

| Variável | Descrição | Valor Padrão |
|----------|-----------|--------------|
| `MONGODB_URL` | URL de conexão do MongoDB | `mongodb://localhost:27017` |
| `MONGODB_DB` | Nome do database | `auctions` |
| `AUCTION_DURATION` | Duração até finalização automática do leilão | `24h` |
| `BATCH_INSERT_INTERVAL` | Intervalo para processamento batch de inserções | `7m` |
| `MAX_BATCH_SIZE` | Tamanho máximo do batch de inserções | `10` |

**Notas:**
- O arquivo `.env` é **opcional** quando rodando via Docker - as variáveis são definidas no `docker-compose.yml`
- Para desenvolvimento local, configure o arquivo `cmd/auction/.env` com as variáveis necessárias
- Quando rodando via Docker Compose, `MONGODB_URL` é sobrescrita automaticamente para `mongodb://auction-mongodb:27017`
- `AUCTION_DURATION` e `BATCH_INSERT_INTERVAL` aceitam unidades: `s` (segundos), `m` (minutos), `h` (horas)
- Ajuste `MAX_BATCH_SIZE` conforme o volume de operações da sua aplicação
- Leilões são finalizados automaticamente após `AUCTION_DURATION` (padrão: 24h)

## 🐛 Troubleshooting

### Porta 8080 já em uso
```bash
# Ver o que está usando a porta
lsof -i :8080

# Parar os containers
make docker-down
```

### MongoDB não conecta
```bash
# Ver logs
make docker-logs

# Verificar health do MongoDB
docker-compose ps

# Resetar tudo
make clean
make run
```

### Erro ao buildar
```bash
# Rebuild sem cache
make docker-build

# Ou manualmente
docker-compose build --no-cache
docker-compose up -d
```

## 📝 Licença

MIT
