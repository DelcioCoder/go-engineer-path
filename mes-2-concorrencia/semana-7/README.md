# Semana 7: Concorrência Avançada e Contextos

> 🎯 **Objetivo:** Proteção em ambientes de produção — cancelamento e timeouts.

---

## 📖 Teoria

### 1. select — Multiplexação de Channels

```go
select {
case msg := <-ch1:
    fmt.Println("recebido de ch1:", msg)
case msg := <-ch2:
    fmt.Println("recebido de ch2:", msg)
case <-time.After(5 * time.Second):
    fmt.Println("timeout!")
default:
    fmt.Println("nenhum canal pronto")
}
```

#### Padrão: Timeout com select
```go
select {
case resultado := <-ch:
    fmt.Println("resultado:", resultado)
case <-time.After(3 * time.Second):
    fmt.Println("demorou demais, cancelando...")
}
```

### 2. Fechando Canais

```go
ch := make(chan int, 5)
ch <- 1
ch <- 2
ch <- 3
close(ch) // Sinaliza: "não vou enviar mais"

// Ler de canal fechado retorna zero value
// Range para automaticamente quando canal fecha
for v := range ch {
    fmt.Println(v) // 1, 2, 3
}

// Verificar se canal foi fechado
val, ok := <-ch
if !ok {
    fmt.Println("canal fechado")
}
```

> ⚠️ **NUNCA** envie para um canal fechado — causa panic!

### 3. context — O Pacote Mais Fundamental para Produção

```go
// Background — contexto raiz (nunca cancela)
ctx := context.Background()

// WithTimeout — cancela automaticamente após duração
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // SEMPRE chame cancel!

// WithCancel — cancelamento manual
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Verificar cancelamento
select {
case <-ctx.Done():
    fmt.Println("cancelado:", ctx.Err())
    return
default:
    // continua trabalhando
}
```

### 4. Propagação de Contexto

```go
// Padrão: context é SEMPRE o primeiro argumento
func processarPedido(ctx context.Context, pedidoID int) error {
    // Cria sub-contexto com timeout menor
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()

    select {
    case <-ctx.Done():
        return ctx.Err() // context.DeadlineExceeded ou context.Canceled
    case resultado := <-fazerTrabalho(ctx):
        return salvar(ctx, resultado)
    }
}
```

### 5. context.WithValue (Use com MUITA moderação)

```go
// Para dados de escopo de request (request ID, user ID, etc.)
type chave string
const chaveRequestID chave = "request_id"

ctx = context.WithValue(ctx, chaveRequestID, "abc-123")

// Recuperar
if id, ok := ctx.Value(chaveRequestID).(string); ok {
    fmt.Println("Request:", id)
}
```

> ⚠️ **NÃO use** context.WithValue para passar parâmetros de função!

---

## 🔨 Prática

**Projeto:** Worker Pool com Timeout Global via Context

### Requisitos:
- [ ] Pegar o Worker Pool da Semana 6
- [ ] Adicionar `context.WithTimeout` de 5 segundos
- [ ] Workers devem verificar `ctx.Done()` antes de processar
- [ ] Se timeout, cancelar **TODAS** as goroutines graciosamente
- [ ] Usar `select` para lidar com múltiplos canais
- [ ] Imprimir quantos jobs completaram vs cancelaram
- [ ] Workers devem simular trabalho lento (alguns > 5s)

### Diagrama de fluxo:
```
                    ┌──── context.WithTimeout ────┐
                    │                              │
  [Jobs] ──▶ Worker Pool ──▶ [Resultados]         │
                    │                              │
                    └──── ctx.Done() = cancelar ──┘
```

---

## 📚 Referências da Semana
- [Go By Example — Select](https://gobyexample.com/select)
- [Go By Example — Timeouts](https://gobyexample.com/timeouts)
- [Go Blog — Contexts and Structs](https://go.dev/blog/context-and-structs)
- [Go Blog — Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [Go Doc — context package](https://pkg.go.dev/context)
