# Semana 12: Projeto Final de Portfólio — Event Processor

> 🎯 **Objetivo:** Juntar tudo das 12 semanas numa aplicação robusta e impressionante.

---

## 📖 Definição do Sistema

### Event Processor / Notificador Assíncrono

Um sistema que recebe milhares de eventos JSON via HTTP, processa-os concorrentemente com worker pools, persiste no banco de dados e tenta enviar notificações — tudo com proteção de contexto, logging estruturado e shutdown gracioso.

### Fluxo do Sistema:

```
┌─────────────┐     ┌──────────────┐     ┌──────────────┐     ┌──────────┐
│  HTTP POST  │────▶│  Validação   │────▶│  Canal Jobs  │────▶│ Worker   │
│  /events    │     │  Rápida      │     │  (buffered)  │     │ Pool     │
│             │     │              │     │              │     │ (N=10)   │
│  JSON Body  │     │  201 Aceito  │     │              │     │          │
└─────────────┘     └──────────────┘     └──────────────┘     └────┬─────┘
                                                                    │
                                                               ┌────▼─────┐
                                                               │ Persistir│
                                                               │ no banco │
                                                               │    +     │
                                                               │ Enviar   │
                                                               │ "email"  │
                                                               │ (fake)   │
                                                               └──────────┘
```

---

## 🔨 Requisitos Técnicos

### De cada semana, o que usar:

| Semana | Conceito | Aplicação no Projeto |
|--------|----------|---------------------|
| 1 | Sintaxe, funções | Base do código |
| 2 | Structs, slices, maps | Modelos de Evento e Resultado |
| 3 | Interfaces, erros | Repository interface, erros customizados |
| 4 | Testes | Table-driven tests, 80%+ cobertura |
| 5 | Goroutines | Workers concorrentes |
| 6 | Channels, WaitGroup | Worker pool com canais |
| 7 | Context | Timeout no processamento |
| 8 | Benchmarks | Profiling da pipeline |
| 9 | net/http, JSON | API REST para receber eventos |
| 10 | database/sql, Clean Arch | Persistência com padrão Repository |
| 11 | slog, Graceful Shutdown | Logs JSON, encerramento seguro |

### Checklist do Projeto:

#### API HTTP
- [ ] Endpoint `POST /events` — recebe JSON, valida, retorna `201 Accepted`
- [ ] Endpoint `GET /events` — listar eventos processados
- [ ] Endpoint `GET /health` — health check
- [ ] Middleware de logging com `slog` (método, rota, status, duração)
- [ ] Middleware de request ID

#### Processamento Concorrente
- [ ] Canal buffered para receber jobs
- [ ] Pool de N workers (configurável via env)
- [ ] Cada worker: persiste no banco + "envia email" (fake)
- [ ] `context.WithTimeout` protege cada job (ex: max 3s)
- [ ] Workers param graciosamente quando contexto é cancelado

#### Persistência
- [ ] PostgreSQL via `database/sql`
- [ ] Tabela `events` com status (pendente, processado, falha)
- [ ] Interface `EventRepository`
- [ ] Pool de conexões configurado

#### Produção
- [ ] Graceful Shutdown (SIGINT/SIGTERM)
- [ ] Logging JSON com `log/slog`
- [ ] Configuração via variáveis de ambiente
- [ ] Testes com mock repository

#### Infraestrutura
- [ ] Dockerfile multistage (builder + alpine/scratch)
- [ ] docker-compose com API + PostgreSQL
- [ ] README espetacular

---

## 📐 Estrutura do Projeto

```
projeto-final/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handler/
│   │   ├── event_handler.go
│   │   └── health_handler.go
│   ├── middleware/
│   │   ├── logging.go
│   │   └── request_id.go
│   ├── model/
│   │   └── event.go
│   ├── repository/
│   │   ├── event_repo.go        (interface)
│   │   ├── postgres_repo.go     (implementação)
│   │   └── mock_repo.go         (para testes)
│   ├── service/
│   │   └── event_service.go
│   └── worker/
│       └── pool.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## 📚 Referências da Semana
- Tudo das 11 semanas anteriores!
- [Docker — Multi-stage Builds](https://docs.docker.com/build/building/multi-stage/)
- [Docker Compose](https://docs.docker.com/compose/)
- [12-Factor App](https://12factor.net/)
