# Load Test CLI

Ferramenta de teste de carga em linha de comando para medir performance de serviços HTTP.

## Uso Rápido com Docker

A forma mais rápida de usar a ferramenta é através da imagem pré-compilada no Docker Hub:

```bash
docker run --rm gbuenodev/loadtest \
  --url=http://example.com \
  --requests=100 \
  --concurrency=10
```

## Instalação Local

### Requisitos

- Go 1.25.5 ou superior
- Docker (opcional, para build local)

### Build a partir do código-fonte

```bash
# Clonar o repositório
git clone <repository-url>
cd desafio07

# Build local
make build
```

## Uso

### CLI Local

```bash
# Executar diretamente
go run ./cmd/loadtest --url=http://example.com --requests=100 --concurrency=10

# Ou usando o binário compilado
./bin/loadtest --url=http://example.com --requests=100 --concurrency=10
```

### Docker (build local)

```bash
# Build da imagem
make docker-build

# Executar
docker run --rm loadtest:latest \
  --url=http://example.com \
  --requests=100 \
  --concurrency=10
```

## Parâmetros

- `--url`: URL do serviço a ser testado (obrigatório)
- `--requests`: Número total de requisições (obrigatório)
- `--concurrency`: Número de requisições simultâneas (obrigatório)

## Relatório

A ferramenta gera um relatório contendo:

- Tempo total de execução
- Total de requisições realizadas
- Quantidade de requisições com status 200
- Distribuição de outros códigos de status HTTP
- Erros de rede/timeout

## Desenvolvimento

```bash
# Executar testes
make test

# Build local
make build

# Executar exemplo local
make run-local
```

## Comandos Make Disponíveis

- `make build` - Compila o binário
- `make run-local` - Executa um exemplo local
- `make docker-build` - Cria a imagem Docker
- `make docker-run` - Executa via Docker
- `make test` - Executa os testes com race detector
