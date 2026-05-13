# Magic Bullet

Monorepo boilerplate com `Go + Gin + Huma` no backend, `Astro` no frontend e sincronizacao de tipos via OpenAPI.

## Estrutura

```text
magic-bullet/
├── backend/
├── frontend/
├── infra/
└── Makefile
```

## Requisitos

- Go 1.25+
- Node 20+
- Docker + Docker Compose

## Primeiros passos

```bash
cp .env.example .env
make up
cd backend && go mod tidy
npm --prefix frontend install
make sync-types
```

## Desenvolvimento

```bash
make dev
```

## Produção com Docker

Crie um arquivo `.env` a partir do exemplo:

```bash
cp .env.example .env
```

Build das imagens:

```bash
make prod-build
```

Subir stack de produção:

```bash
make prod-up
```

Logs:

```bash
make prod-logs
```

Encerrar:

```bash
make prod-down
```

## Fluxo de tipos

1. Atualize uma operacao ou schema no backend.
2. Rode `make sync-types`.
3. O backend gera `backend/openapi/openapi.yaml`.
4. O frontend atualiza `frontend/src/api/schema.d.ts`.

## Rotas base

- API: `http://localhost:8080/api/v1`
- Docs: `http://localhost:8080/api/v1/docs`
- OpenAPI 3.1: `http://localhost:8080/api/v1/openapi.json`
- OpenAPI 3.0.3: `http://localhost:8080/api/v1/openapi-3.0.yaml`
- Health: `http://localhost:8080/health`
- Frontend SSR: `http://localhost:4321`

## Frontend UI

- Tailwind CSS v4 configurado via `@tailwindcss/vite`
- React habilitado via `@astrojs/react`
- Tema base aplicado em `frontend/src/styles/globals.css`
- Fonte padrão definida como `Inter`
- shadcn/ui preparado com:
  - `frontend/components.json`
  - `frontend/src/lib/utils.ts`
  - `frontend/src/components/ui/button.tsx`
