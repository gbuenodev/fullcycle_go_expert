# Weather App - Frontend

Frontend React + TypeScript para consultar informações de clima baseado no CEP brasileiro.

## Tecnologias

- **React 18** - Biblioteca UI
- **TypeScript** - Tipagem estática
- **Vite** - Build tool e dev server
- **Material-UI (MUI)** - Componentes UI
- **Axios** - Cliente HTTP
- **Zod** - Validação de esquemas
- **React Input Mask** - Máscara de input

## Arquitetura

```
src/
├── api/              # Configuração Axios
├── components/       # Componentes React
│   ├── Loading.tsx
│   ├── WeatherForm.tsx
│   └── WeatherResult.tsx
├── context/          # Context API
│   ├── WeatherContext.tsx
│   └── WeatherProvider.tsx
├── hooks/            # Custom hooks
│   └── useWeather.ts
├── services/         # Serviços de API
│   └── weatherService.ts
└── utils/            # Utilitários
    └── validation.ts
```

## Funcionalidades

- Formulário com validação de CEP brasileiro
- Máscara automática (formato: 00000-000)
- Consulta de temperatura em Celsius, Fahrenheit e Kelvin
- Tratamento de erros (404, 422)
- Loading states
- UI responsiva e moderna

## Como usar

## Desenvolvimento Local

### Pré-requisitos

- Node.js 18+
- npm ou yarn

### Instalação

```bash
# Instalar dependências
npm install

# Configurar variáveis de ambiente
cp .env.example .env
# Editar .env com a URL do backend
```

### Executar

```bash
# Modo desenvolvimento
npm run dev

# Build de produção
npm run build

# Preview do build
npm run preview

# Lint
npm run lint
```

## Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
VITE_BACKEND_URL=http://localhost:3000
```

## Docker

### Build da imagem

```bash
docker build -t weather-app .
```

### Executar localmente

```bash
docker run -p 8080:8080 weather-app
```

Acesse: http://localhost:8080

## Deploy no Google Cloud Run

### Pré-requisitos

- Google Cloud CLI instalado
- Projeto GCP criado
- Billing habilitado

### Deploy com Makefile

```bash
# 1. Configurar projeto e habilitar APIs (uma vez)
make setup-gcloud PROJECT_ID=seu-projeto-id

# 2. Deploy
make deploy BACKEND_URL=https://seu-backend.run.app

# 3. Ver URL do serviço
make url

# Atualizar variável de ambiente depois
make update-env BACKEND_URL=https://novo-backend.run.app
```

## API Backend

O frontend espera um backend com o seguinte endpoint:

```
GET /weather/{cep}
```

**Respostas:**

- `200 OK`:
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

- `404 Not Found`: CEP não encontrado
- `422 Unprocessable Entity`: CEP inválido

## Comandos Makefile

| Comando | Descrição |
|---------|-----------|
| `make help` | Mostra todos os comandos |
| `make install` | Instala dependências |
| `make dev` | Servidor de desenvolvimento |
| `make build` | Build de produção |
| `make lint` | Executa linter |
| `make docker-build` | Build da imagem Docker |
| `make docker-run` | Roda container localmente |
| `make deploy` | Deploy no Cloud Run |

## Licença

MIT
