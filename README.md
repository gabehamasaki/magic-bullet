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
make up
cd backend && go mod tidy
npm --prefix frontend install
make sync-types
```

## Desenvolvimento

Terminal 1:

```bash
make dev-backend
```

Terminal 2:

```bash
make dev-frontend
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
