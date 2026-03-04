# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copiar dependências primeiro (cache de build)
COPY go.mod go.sum ./
RUN go mod download

# Copiar código-fonte
COPY . .

# Compilar binário estático
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/api

# ---

# Runtime stage (imagem final super leve)
FROM alpine:3.19

# Certificados SSL (para HTTPS externo)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copiar apenas o binário
COPY --from=builder /app/server .

EXPOSE 8080

# Executar como não-root
RUN adduser -D -g '' appuser
USER appuser

ENTRYPOINT ["./server"]
