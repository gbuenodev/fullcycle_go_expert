# Weather App - Frontend

Interface web moderna para consulta de clima por CEP brasileiro.

## Stack

- React 18
- TypeScript
- Vite (build tool)
- Material-UI (MUI)
- Axios (HTTP client)
- Zod (validation)
- React Input Mask

## Estrutura

```
frontend/
├── src/
│   ├── components/       # Componentes React
│   │   ├── Loading.tsx
│   │   ├── WeatherForm.tsx
│   │   └── WeatherResult.tsx
│   ├── context/          # Context API
│   │   ├── WeatherContext.tsx
│   │   └── WeatherProvider.tsx
│   ├── hooks/            # Custom hooks
│   │   └── useWeather.ts
│   ├── services/         # API calls
│   │   └── weatherService.ts
│   ├── theme/            # MUI theme config
│   ├── utils/            # Utilitários
│   │   └── validation.ts
│   ├── api.ts            # Axios instance
│   └── main.tsx          # Entry point
└── package.json
```

## Funcionalidades

- Formulário com validação de CEP brasileiro
- Máscara automática (formato: 00000-000)
- Consulta de temperatura em Celsius, Fahrenheit e Kelvin
- Exibição da cidade encontrada
- Tratamento de erros (400, 404, 422, 500)
- Loading states
- UI responsiva e moderna

## Configuração

### Environment Variables

**Build time:**
```env
VITE_BACKEND_URL=http://localhost:3000
```

Esta variável é injetada durante o build via Docker build args.

### Docker Build

O Dockerfile aceita build arg:
```bash
docker build --build-arg VITE_BACKEND_URL=http://localhost:3000 -t frontend .
```

## Desenvolvimento Local

### Pré-requisitos

- Node.js 18+
- npm

### Instalação

```bash
# Instalar dependências
npm install

# Dev server (hot reload)
npm run dev

# Build de produção
npm run build

# Preview do build
npm run preview

# Linter
npm run lint
```

## Comandos Make

```bash
make help          # Lista todos os comandos
make install       # npm install
make dev           # Inicia dev server
make build         # Build de produção
make lint          # Executa linter
make docker-build  # Build da imagem Docker
make docker-run    # Roda container
```

## Integração com Backend

**API Service:** `src/services/weatherService.ts`

```typescript
const { data } = await api.post<WeatherResponse>('/weather', { cep });
```

**Request:**
```json
{
  "cep": "01310100"
}
```

**Response type:**
```typescript
interface WeatherResponse {
  city: string;
  temp_C: number;
  temp_F: number;
  temp_K: number;
}
```

**Error handling:**
- 400: Request body inválido
- 422: CEP inválido
- 404: CEP não encontrado
- 500: Erro interno

## UI Components

### WeatherForm
- Input field com máscara de CEP
- Validação com Zod
- Submit handler
- Estados de loading

### WeatherResult
- Display de temperatura (C, F, K)
- Display de cidade
- Estados de loading/error
- Formatação de números

### Loading
- Spinner de carregamento
- Feedback visual durante requests

### Theme
- MUI theme customizado
- Cores e tipografia consistentes
- Componentes estilizados

## Build & Deploy

O frontend é servido via Nginx no container:

**Port:** 8080
**Health check:** `GET /health`

## Acesso

**Development:** http://localhost:5173 (Vite dev server)
**Production (Docker):** http://localhost:8080
