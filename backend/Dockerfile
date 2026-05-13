FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY backend/go.mod backend/go.sum ./backend/

WORKDIR /src/backend
RUN go mod download

COPY backend ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

FROM alpine:3.22 AS runtime

RUN apk add --no-cache ca-certificates tzdata && \
    addgroup -S app && \
    adduser -S app -G app

WORKDIR /app

COPY --from=builder /out/api /app/api

ENV GIN_MODE=release \
    HTTP_ADDR=:8080

USER app

EXPOSE 8080

CMD ["./api"]
