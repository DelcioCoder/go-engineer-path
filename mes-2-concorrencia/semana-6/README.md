# Semana 6: Sincronização e Canais

> 🎯 **Objetivo:** Aprender a compartilhar memória comunicando-se.

---

## 📖 Teoria

### 1. sync.WaitGroup

```go
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)          // Incrementa ANTES de lançar
    go func(id int) {
        defer wg.Done() // Decrementa ao terminar
        fmt.Println("worker", id)
    }(i)               // Passa i como argumento (evita closure bug!)
}

wg.Wait() // Bloqueia até todas as goroutines terminarem
```

### 2. sync.Mutex e sync.RWMutex

```go
// Mutex — acesso exclusivo (leitura E escrita bloqueiam)
var mu sync.Mutex
var contador int

go func() {
    mu.Lock()
    contador++
    mu.Unlock()
}()

// RWMutex — múltiplas leituras, uma escrita
var rw sync.RWMutex

// Leitura (múltiplos leitores simultâneos)
rw.RLock()
valor := contador
rw.RUnlock()

// Escrita (exclusiva)
rw.Lock()
contador++
rw.Unlock()
```

### 3. Detector de Data Race

```bash
# SEMPRE use durante desenvolvimento!
go run -race main.go
go test -race ./...
```

```go
// Exemplo de data race (Bug SÉRIO)
var x int
go func() { x++ }()  // goroutine 1
go func() { x++ }()  // goroutine 2
// Race condition! Resultado imprevisível.
```

### 4. Channels (Canais)

```go
// Unbuffered (sem buffer) — síncrono
ch := make(chan int)       // Bloqueia até alguém ler
ch <- 42                   // Enviar (bloqueia)
valor := <-ch              // Receber (bloqueia)

// Buffered (com buffer) — assíncrono até encher
ch := make(chan int, 10)   // Capacidade de 10
ch <- 42                   // Não bloqueia (buffer não está cheio)
```

### 5. Direcionamento de Canais

```go
// Canal bidirecional (padrão)
func processar(ch chan int)

// Canal apenas de envio (send-only)
func produzir(ch chan<- int) {
    ch <- 42  // OK
    // <-ch   // ERRO de compilação!
}

// Canal apenas de recepção (receive-only)
func consumir(ch <-chan int) {
    v := <-ch  // OK
    // ch <- 1 // ERRO de compilação!
}
```

---

## 🔨 Prática

### Parte 1: Consertar o download da Semana 5

```go
var wg sync.WaitGroup

for _, url := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        downloadURL(u)
    }(url)
}

wg.Wait() // Agora espera todos terminarem!
```

### Parte 2: Worker Pool com Channels

```
     ┌─────────┐     ┌──────────┐     ┌───────────┐
     │  Jobs   │────▶│ Workers  │────▶│ Resultados│
     │ Channel │     │ (N goroutines) │  Channel  │
     └─────────┘     └──────────┘     └───────────┘
```

```go
func worker(id int, jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        resultado := processar(job)
        results <- resultado
    }
}
```

### Requisitos:
- [ ] Consertar download da Semana 5 com `sync.WaitGroup`
- [ ] Medir tempo e comparar síncrono vs concorrente
- [ ] Usar `go run -race` para verificar data races
- [ ] Criar Worker Pool com canais de jobs e resultados
- [ ] Limitar número de workers (ex: 5 simultâneos)
- [ ] Demonstrar canal unbuffered vs buffered

---

## 📚 Referências da Semana
- [Go By Example — WaitGroups](https://gobyexample.com/waitgroups)
- [Go By Example — Mutexes](https://gobyexample.com/mutexes)
- [Go By Example — Channels](https://gobyexample.com/channels)
- [Go By Example — Channel Directions](https://gobyexample.com/channel-directions)
- [Go Blog — Share Memory By Communicating](https://go.dev/blog/codelab-share)
- [Go Blog — Race Detector](https://go.dev/doc/articles/race_detector)
