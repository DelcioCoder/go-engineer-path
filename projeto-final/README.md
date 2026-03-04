# 🚀 Event Processor — Projeto Final Go Engineer Path

> Um processador de eventos assíncrono que demonstra domínio de Go: concorrência com worker pools, API REST pura, persistência PostgreSQL, logging estruturado e deploy via Docker.

---

## 📐 Arquitetura

```
                    ┌─────────────────────────────────────────┐
                    │           Event Processor API           │
                    │                                         │
  HTTP POST ──────▶│  Handler → Validação → 201 Accepted    │
  /events          │              │                          │
                    │              ▼                          │
                    │        Canal de Jobs                    │
                    │         (buffered)                      │
                    │              │                          │
                    │    ┌────────┼────────┐                 │
                    │    ▼        ▼        ▼                 │
                    │  Worker  Worker  Worker  (N workers)   │
                    │    │        │        │                  │
                    │    ▼        ▼        ▼                 │
                    │  ┌──────────────────────┐              │
                    │  │  PostgreSQL + Email   │              │
                    │  └──────────────────────┘              │
                    └─────────────────────────────────────────┘
```

## 🛠️ Tecnologias

| Tecnologia | Uso |
|------------|-----|
| Go 1.22+ | Standard library pura (net/http, database/sql, log/slog) |
| PostgreSQL | Persistência de eventos |
| Docker | Multistage build (imagem final < 20MB) |
| docker-compose | Orquestração local |

## 🚀 Como Executar

### Com Docker (recomendado)
```bash
docker-compose up --build
```

### Localmente
```bash
# Variáveis de ambiente
export PORT=8080
export DATABASE_URL=postgres://user:pass@localhost:5432/events?sslmode=disable
export WORKER_COUNT=10

# Executar
go run ./cmd/api
```

## 📡 Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| `POST` | `/events` | Enviar evento para processamento |
| `GET` | `/events` | Listar eventos processados |
| `GET` | `/health` | Health check |

### Exemplo de uso:
```bash
# Enviar evento
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{"tipo": "campanha", "destinatario": "user@email.com", "dados": {"titulo": "Promoção"}}'

# Listar eventos
curl http://localhost:8080/events

# Health check
curl http://localhost:8080/health
```

## 🧪 Testes

```bash
go test -v -race -cover ./...
```

## 📊 Benchmarks

```bash
go test -bench=. -benchmem ./...
```

## 🏗️ Estrutura do Projeto

```
projeto-final/
├── cmd/api/main.go              ← Ponto de entrada
├── internal/
│   ├── handler/                 ← Handlers HTTP
│   ├── middleware/               ← Logging, Request ID
│   ├── model/                   ← Structs de domínio
│   ├── repository/              ← Interface + Postgres + Mock
│   ├── service/                 ← Lógica de negócio
│   └── worker/                  ← Worker pool
├── Dockerfile                   ← Multistage build
├── docker-compose.yml           ← API + PostgreSQL
├── go.mod
└── README.md                    ← Você está aqui
```

## 📝 Decisões de Engenharia

### Por que Channels e não Mutex?
> *"Don't communicate by sharing memory; share memory by communicating."* — Go Proverb

Canais modelam naturalmente o fluxo produtor→consumidor do sistema. Cada evento é um job que flui pelo canal até um worker disponível, sem necessidade de locks explícitos.

### Por que Standard Library e não frameworks?
A standard library do Go é excepcionalmente completa. `net/http` com o novo `ServeMux` (Go 1.22+) elimina a necessidade de frameworks para 90% dos casos.

### Métricas de Performance
| Métrica | Valor |
|---------|-------|
| Latência P99 | TODO |
| Throughput (req/s) | TODO |
| Imagem Docker | TODO |

---

> Construído durante o **Go Engineer Path** — 12 semanas de estudo intensivo.
