# Semana 11: Operações de Produção

> 🎯 **Objetivo:** O que a maioria esquece de fazer — preparar para o mundo real.

---

## 📖 Teoria

### 1. Tratamento de Sinais do Sistema Operacional

```go
import (
    "os"
    "os/signal"
    "syscall"
)

// Capturar SIGINT (Ctrl+C) e SIGTERM (docker stop)
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

<-quit // Bloqueia até receber sinal
log.Println("Sinal recebido, encerrando...")
```

### 2. Graceful Shutdown

```go
func main() {
    srv := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    // Goroutine para servir
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Erro no servidor: %v", err)
        }
    }()

    log.Println("Servidor rodando em :8080")

    // Esperar sinal de shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Desligando servidor...")

    // Contexto com timeout para shutdown gracioso
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Pára de aceitar novos requests e espera os em andamento
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Shutdown forçado: %v", err)
    }

    // Fechar banco, flush logs, etc.
    db.Close()

    log.Println("Servidor encerrado com sucesso ✓")
}
```

### 3. Logging Estruturado com log/slog (Go 1.21+)

```go
import "log/slog"

// Logger padrão (texto)
slog.Info("servidor iniciado", "porta", 8080)
// 2024/01/01 10:00:00 INFO servidor iniciado porta=8080

// Logger JSON (para produção)
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo,
}))
slog.SetDefault(logger)

slog.Info("request processado",
    "method", "GET",
    "path", "/todos",
    "status", 200,
    "duracao_ms", 12,
)
// {"time":"2024-01-01T10:00:00Z","level":"INFO","msg":"request processado","method":"GET","path":"/todos","status":200,"duracao_ms":12}

// Logger com contexto (adiciona campos a todos os logs)
logger = logger.With("servico", "todo-api", "versao", "1.0.0")

// Grupos para organizar
slog.Info("conexão db",
    slog.Group("db",
        slog.String("host", "localhost"),
        slog.Int("porta", 5432),
    ),
)

// Níveis: Debug, Info, Warn, Error
slog.Debug("detalhe interno")     // Só aparece se Level=Debug
slog.Warn("recurso quase esgotado", "uso_pct", 85)
slog.Error("falha na query", "err", err)
```

### 4. Variáveis de Ambiente

```go
// Nativo — sem bibliotecas externas!
port := os.Getenv("PORT")
if port == "" {
    port = "8080"  // Valor padrão
}

dbURL := os.Getenv("DATABASE_URL")
if dbURL == "" {
    log.Fatal("DATABASE_URL não configurada")
}

// Helper simples
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
```

### 5. Middleware de Logging com slog

```go
func loggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            inicio := time.Now()

            // Wrapper para capturar status code
            ww := &responseWriter{ResponseWriter: w, status: http.StatusOK}
            next.ServeHTTP(ww, r)

            logger.Info("request",
                "method", r.Method,
                "path", r.URL.Path,
                "status", ww.status,
                "duracao_ms", time.Since(inicio).Milliseconds(),
                "remote_addr", r.RemoteAddr,
            )
        })
    }
}

type responseWriter struct {
    http.ResponseWriter
    status int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.status = code
    rw.ResponseWriter.WriteHeader(code)
}
```

---

## 🔨 Prática

**Projeto:** Adicionar produção-readiness à API REST

### Requisitos:
- [ ] Implementar Graceful Shutdown (capturar SIGINT/SIGTERM)
- [ ] Esperar requests em andamento terminarem antes de encerrar
- [ ] Substituir `log` por `log/slog` em toda a API
- [ ] Logs em formato JSON direcionados para `os.Stdout`
- [ ] Middleware de logging com: método, rota, status, duração
- [ ] Configuração via variáveis de ambiente (`PORT`, `DATABASE_URL`)
- [ ] Função `getEnv` com valores padrão
- [ ] Testar: iniciar servidor, fazer request, Ctrl+C, verificar shutdown limpo

---

## 📚 Referências da Semana
- [Go Doc — log/slog](https://pkg.go.dev/log/slog)
- [Go Doc — os/signal](https://pkg.go.dev/os/signal)
- [Go Blog — Structured Logging with slog](https://go.dev/blog/slog)
- [Go By Example — Signals](https://gobyexample.com/signals)
- [Go By Example — Environment Variables](https://gobyexample.com/environment-variables)
