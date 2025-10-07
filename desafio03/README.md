# Desafio 03 — **Listagem de Orders** (HTTP, GraphQL, gRPC com Evans)

> **Foco:** este guia cobre **apenas** a rota de **ListOrders** nas três interfaces do projeto: **HTTP/REST**, **GraphQL** e **gRPC** (usando **Evans** como client).

Repositório original: [`goexpert/20-CleanArch`](https://github.com/devfullcycle/goexpert/tree/main/20-CleanArch)

---

## ✅ Pré-requisitos

- **Go 1.25+** instalado
- **Docker** e **Docker Compose** instalados *(caso opte por MySQL)*
- **Evans** (cliente gRPC em modo REPL)

---

## 📦 Clonar o projeto

```bash
git clone https://github.com/gbuenodev/fullcycle_go_expert.git
cd fullcycle_go_expert/desafio03
```

---

## ⚙️ Configuração do `.env`

Crie um arquivo `.env` em `cmd/ordersystem/`. Escolha **uma** das opções abaixo.

### Opção A — Usando **SQLite** (mais simples)

```env
# Banco
DB_DRIVER=sqlite
DB_DSN=./data/orders.db

# Portas dos servidores
HTTP_PORT=8080
GRAPHQL_PORT=8081
GRPC_PORT=50051
```

### Opção B — Usando **MySQL** (via Docker)

```env
# Banco
DB_DRIVER=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASS=root
DB_NAME=orders
DB_DSN=${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true

# Portas dos servidores
HTTP_PORT=8080
GRAPHQL_PORT=8081
GRPC_PORT=50051
```

> Ajuste as variáveis de acordo com seu `docker-compose.yaml` se necessário.

---

## 🐳 Subir infraestrutura (apenas para MySQL)

Se optou pela **Opção B (MySQL)**:

```bash
make docker-up
```

---

## 🧱 Migrações

As migrações estão em `./migrations`.

```bash
make migrate-up
```

---

## 🚀 Iniciar a aplicação

A aplicação principal fica em `cmd/ordersystem`:

```bash
make server-up
```

Você deverá ver logs indicando que os servidores **HTTP**, **GraphQL** e **gRPC** estão rodando nas portas definidas no `.env` (por padrão: `8080`, `8081`, `50051`).

---

# 🔎 ListOrders por API

Como executar **ListOrders**.

> Se a resposta vier vazia (`[]`), crie/insira pedidos antes de listar.

---

## A) HTTP/REST — `ListOrders`

**Endpoint (sugerido):** `GET http://localhost:${HTTP_PORT}/order`

```bash
curl -s http://localhost:8080/order | jq .
```

**Resposta esperada (exemplo):**
```json
[
  {
    "id": "01J...XYZ",
    "Price": 120.5,
    "Tax": 12.05,
    "FinalPrice": 132.55,
  }
]
```

---

## B) GraphQL — `ListOrders`

**Endpoint:** `http://localhost:${GRAPHQL_PORT}/query`  

**Query (listar todos):**
```graphql
query queryOrders{
  orders {
    id
    Price
    FinalPrice
    Tax
  }
}
```

**Resposta esperada (exemplo):**
```json
{
  "data": {
    "orders": [
      {
        "id": "abc",
        "Price": 100,
        "FinalPrice": 112,
        "Tax": 12
      },
      {
        "id": "def",
        "Price": 100.5,
        "FinalPrice": 101,
        "Tax": 0.5
      },
      {
        "id": "ghi",
        "Price": 79,
        "FinalPrice": 202.3,
        "Tax": 123.3
      }
    ]
  }
}
```

---

## C) gRPC — `ListOrders` usando **Evans**

Os arquivos `.proto` estão em:  
`internal/infra/grpc/protofiles/order.proto`  
Os stubs gerados ficam em:  
`internal/infra/grpc/pb/`

### 1) Usando **Reflection** (se habilitada no servidor gRPC)

Abra o REPL do Evans apontando para o host/porta:

```bash
evans --host localhost --port 50051 -r repl
```

No REPL do Evans:

```
show package           # opcional: ver pacotes disponíveis
package pb             # selecione o pacote (ajuste caso seu pacote seja diferente)
show service           # opcional: ver serviços
service OrderService   # selecione o serviço correto
call ListOrders        # execute a RPC (envio vazio se o request não tem campos)
{}
```

**Resposta (exemplo):**
```json
{
  "orders": [
    { "id": "01J...XYZ", "Price": 120.5, "Tax": 12.05, "FinalPrice": 132.55 }
  ]
}
```

### 2) **Sem Reflection** (carregando o `.proto` manualmente)

Se o servidor **não** tiver reflection habilitada, informe o pacote e o arquivo `.proto`:

```bash
evans --host localhost --port 50051 \
  -p internal/infra/grpc/protofiles \
  -f internal/infra/grpc/protofiles/order.proto \
  repl
```

No REPL:

```
show package
package pb                   # ou o nome do pacote definido no .proto
show service
service OrderService
call ListOrders
{}
```

---

## 🛠️ Troubleshooting

- **Portas diferentes** das usadas no README? Ajuste `HTTP_PORT`, `GRAPHQL_PORT`, `GRPC_PORT` no `.env` e reinicie a aplicação.
- **Erro com MySQL**: confirme que `docker compose up -d` está ativo e que `DB_DSN` bate com suas credenciais/DB.
- **Lista vazia**: insira pedidos (seed/endpoint/SQL) antes de chamar a listagem.
- **Evans não encontra serviço/método**:
  - Verifique o nome do **pacote** e **serviço** com `show package` / `show service`.
  - Se **sem reflection**, use `-p` (path) e `-f` (proto file) apontando para `internal/infra/grpc/protofiles/order.proto`.
- **Erros de import Go**: rode `go mod tidy` na raiz e garanta que está executando `go run ./cmd/ordersystem` no módulo correto.

---

## 📂 Estrutura relevante (resumo)

```
desafio03/
├─ cmd/ordersystem/             # Main
├─ internal/
│  ├─ infra/
│  │  ├─ web/                   # HTTP
│  │  ├─ graph/                 # GraphQL
│  │  └─ grpc/
│  │     ├─ pb/                 # Stubs gerados
│  │     ├─ protofiles/         # .proto (order.proto)
│  │     └─ service/            # Implementação gRPC
│  └─ entity/                   # Entidades de domínio (Order, etc.)
├─ migrations/                  # Migrações SQL
├─ api/                         # Exemplos HTTP (.http)
├─ gqlgen.yml                   # Config gqlgen (GraphQL)
└─ .env                         # Configurações (você cria)
```

---

**Pronto!** Com isso você sobe o projeto, configura o ambiente e executa a **listagem de orders** em **HTTP**, **GraphQL** e **gRPC (Evans)**.
