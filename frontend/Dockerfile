FROM golang:1.25-alpine AS openapi

WORKDIR /src/backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend ./

RUN mkdir -p /out && go run ./cmd/api openapi > /out/openapi.yaml

FROM node:24-alpine AS builder

WORKDIR /src

COPY frontend/package.json frontend/package-lock.json ./frontend/

WORKDIR /src/frontend
RUN npm ci

WORKDIR /src
COPY frontend ./frontend
COPY --from=openapi /out/openapi.yaml ./backend/openapi/openapi.yaml

WORKDIR /src/frontend
RUN npm run generate:types && npm run build && npm prune --omit=dev

FROM node:24-alpine AS runtime

WORKDIR /app

COPY --from=builder /src/frontend/package.json /app/package.json
COPY --from=builder /src/frontend/package-lock.json /app/package-lock.json
COPY --from=builder /src/frontend/dist /app/dist

RUN npm ci --omit=dev && \
    npm cache clean --force && \
    addgroup -S app && \
    adduser -S app -G app && \
    chown -R app:app /app

ENV NODE_ENV=production \
    HOST=0.0.0.0 \
    PORT=4321

USER app

EXPOSE 4321

CMD ["node", "./dist/server/entry.mjs"]
