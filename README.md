# Full Cycle Go Expert

Repositório dos desafios e projetos práticos do curso **Go Expert** da Full Cycle.

## Sobre o Curso

O **Go Expert** é um curso avançado de Go focado em aplicações reais, arquitetura de software e boas práticas para desenvolvimento de sistemas distribuídos e APIs de alta performance. O curso aborda desde fundamentos da linguagem até tópicos avançados como observabilidade, clean architecture e protocolos de comunicação.

## Tecnologias Utilizadas

- **Linguagem:** Go 1.23+
- **Bancos de Dados:** SQLite, MySQL, MongoDB
- **Protocolos:** HTTP/REST, GraphQL, gRPC
- **Arquitetura:** Clean Architecture, Dependency Injection (Wire)
- **Observabilidade:** OpenTelemetry, Zipkin, Prometheus
- **Frontend:** React, TypeScript, Material-UI, Vite
- **Infraestrutura:** Docker, Docker Compose
- **Ferramentas:** Evans (gRPC client), Makefile

## Conceitos Estudados

- **Concorrência e Paralelismo** - Goroutines, channels, context
- **APIs Modernas** - REST, GraphQL, gRPC com Protocol Buffers
- **Clean Architecture** - Separação de camadas, inversão de dependências
- **Distributed Tracing** - Rastreamento distribuído com OpenTelemetry
- **Observabilidade** - Métricas, logs e traces
- **Microserviços** - Comunicação entre serviços, propagação de contexto
- **Injeção de Dependências** - Wire (compile-time DI)
- **Database Migrations** - Versionamento de schemas
- **Containerização** - Docker multi-stage builds, Docker Compose

---

## Projetos

### [desafio01](./desafio01) - Cliente-Servidor com gRPC/HTTP

Sistema cliente-servidor utilizando SQLite e migrations.

**Stack:** Go, SQLite, migrations

---

### [desafio02](./desafio02) - Concorrência com APIs Externas

Exercício de concorrência fazendo race entre múltiplas APIs de consulta de CEP.

**Stack:** Go, HTTP clients, context, goroutines

---

### [desafio03](./desafio03) - Sistema de Orders (Multi-Protocol API)

Sistema completo de gerenciamento de pedidos com suporte a três protocolos de comunicação.

**Stack:** Go, Clean Architecture, HTTP/REST, GraphQL, gRPC, Evans, MySQL/SQLite

**[📖 Ver documentação completa](./desafio03/README.md)**

---

### [desafio04](./desafio04) - Weather App com Observabilidade Distribuída

Aplicação fullstack de consulta de clima por CEP com distributed tracing e monitoramento.

**Stack:** Go, React, TypeScript, Material-UI, OpenTelemetry, Zipkin, Prometheus, Docker Compose

**Features:**
- Frontend moderno com React + MUI
- Backend com Clean Architecture
- Distributed Tracing completo
- Métricas com Prometheus
- Input-service para demonstração de tracing entre microserviços

**[📖 Ver documentação completa](./desafio04/README.md)**

**Demo ao vivo:**
- [Frontend](https://weather-app-watebi5u2q-uc.a.run.app/)
- [Backend API](https://weather-api-watebi5u2q-uc.a.run.app/)

---

### [desafio05](./desafio05) - Sistema de Leilões

Sistema de leilões com Clean Architecture e MongoDB.

**Stack:** Go, MongoDB, Clean Architecture, Docker Compose

**Features em desenvolvimento:**
- Entidades de domínio (Auction, Bid)
- Validações de negócio
- Repository pattern
- RESTful API

---

## Como Usar

Cada projeto possui sua própria estrutura e instruções. Navegue até a pasta do desafio e siga o README correspondente (quando disponível).

### Pré-requisitos Gerais

- **Go 1.23+**
- **Docker** e **Docker Compose**
- **Make** (opcional, mas recomendado)

### Quick Start (exemplo geral)

```bash
# Navegar até o projeto desejado
cd desafio03

# Ver comandos disponíveis (se houver Makefile)
make help

# Subir ambiente (se houver Docker)
make docker-up

# Executar aplicação
make run
```

---

## Estrutura do Repositório

```
fullcycle_go_expert/
├── desafio01/          # Cliente-servidor
├── desafio02/          # Concorrência com APIs
├── desafio03/          # Orders (HTTP, GraphQL, gRPC)
├── desafio04/          # Weather App + Observabilidade
├── desafio05/          # Sistema de Leilões
├── graphql_example/    # Exemplos GraphQL
├── grpc_example/       # Exemplos gRPC
└── README.md          # Este arquivo
```

---

## Recursos Adicionais

- [Full Cycle](https://fullcycle.com.br/)
- [Go Documentation](https://go.dev/doc/)
- [gRPC Go Quick Start](https://grpc.io/docs/languages/go/quickstart/)
- [GraphQL Go](https://gqlgen.com/)
- [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/)

---

## Licença

MIT
