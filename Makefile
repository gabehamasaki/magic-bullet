.PHONY: up down install-tools sync-types dev-backend dev-frontend build-backend build-frontend test-backend dev prod-build prod-up prod-down prod-logs

up:
	docker compose --env-file .env -f infra/docker-compose.yml up -d

down:
	docker compose --env-file .env -f infra/docker-compose.yml down

install-tools:
	go install github.com/air-verse/air@latest

install-dependencies:
	cd backend && go mod tidy
	npm --prefix frontend install

sync-types:
	cd backend && go run ./cmd/api openapi > openapi/openapi.yaml
	npm --prefix frontend run generate:types

dev-backend:
	cd backend && air -c .air.toml

dev-frontend:
	npm --prefix frontend run dev

build-backend:
	cd backend && go build ./...

build-frontend:
	npm --prefix frontend run build

test-backend:
	cd backend && go test ./...

dev: up sync-types
	npm --prefix frontend exec -- concurrently --kill-others-on failure --names backend,frontend --prefix name --prefix-colors auto "make dev-backend" "make dev-frontend"

prod-build:
	docker compose --env-file .env -f infra/docker-compose.prod.yml build

prod-up:
	docker compose --env-file .env -f infra/docker-compose.prod.yml up -d --build

prod-down:
	docker compose --env-file .env -f infra/docker-compose.prod.yml down

prod-logs:
	docker compose --env-file .env -f infra/docker-compose.prod.yml logs -f
