# Semana 5: Memória e Fundamentos de Concorrência

> 🎯 **Objetivo:** Entender o que acontece no hardware e dar os primeiros passos com goroutines.

---

## 📖 Teoria

### 1. Stack vs Heap

| | Stack (Pilha) | Heap |
|--|---------------|------|
| **Alocação** | Automática (rápida) | Gerenciada pelo GC (mais lenta) |
| **Escopo** | Local à função | Escapa da função |
| **Tamanho** | Cresce dinamicamente (Go!) | Compartilhado entre goroutines |

```go
// Stack — não escapa
func soma(a, b int) int {
    resultado := a + b  // vive na stack
    return resultado
}

// Heap — escapa!
func novaTarefa(titulo string) *Tarefa {
    t := Tarefa{Titulo: titulo}  // escapa para o heap (retorna ponteiro)
    return &t
}
```

### 2. Escape Analysis

```bash
# Ver decisões do compilador sobre stack/heap
go build -gcflags="-m" ./...

# Saída exemplo:
# ./main.go:10:2: moved to heap: t
# ./main.go:15:6: "hello" does not escape
```

> **Regra:** Se o compilador pode provar que a variável não é usada fora da função, fica na stack. Senão, escapa para o heap.

### 3. Garbage Collector (GC)

- Go tem GC automático — não precisa de `free()` ou `malloc()`
- O GC é **concorrente** (não para 100% do programa)
- Otimizado para **baixa latência** (pausas curtas)
- Variável de ambiente: `GOGC=100` (padrão: coleta quando heap dobra)

### 4. Goroutines — Introdução

```go
// Uma goroutine é uma "thread leve" gerenciada pelo Go
// Custo: ~2KB de stack (vs ~1MB de uma OS thread)

go fazerAlgo()           // Lança goroutine
go func() {             // Goroutine anônima
    fmt.Println("olá")
}()
```

#### Scheduler M:N
- **G** (Goroutine) — unidade de trabalho
- **M** (Machine/OS Thread) — thread do sistema
- **P** (Processor) — contexto de execução (= `GOMAXPROCS`)
- Go mapeia N goroutines em M threads

### 5. ⚠️ Problema Clássico: Programa Termina Cedo

```go
func main() {
    for i := 0; i < 10; i++ {
        go fmt.Println(i)  // Lança goroutine
    }
    // main() termina ANTES das goroutines!
    // Nenhuma saída (ou parcial)
}
```

---

## 🔨 Prática

**Projeto:** Download Síncrono → Concorrente (sem controle)

### Fase 1: Síncrono
```go
// Fazer download de N URLs sequencialmente
// Medir tempo total com time.Since
urls := []string{
    "https://go.dev",
    "https://pkg.go.dev",
    // ... mais URLs
}

inicio := time.Now()
for _, url := range urls {
    downloadURL(url)
}
fmt.Printf("Total: %v\n", time.Since(inicio))
```

### Fase 2: Concorrente (BUGADO de propósito!)
```go
inicio := time.Now()
for _, url := range urls {
    go downloadURL(url)  // Lança goroutine
}
fmt.Printf("Total: %v\n", time.Since(inicio))
// PROBLEMA: main termina antes dos downloads!
```

### Requisitos:
- [ ] Criar programa que faz download real (HTTP GET) de 10+ URLs
- [ ] Versão síncrona funcionando e medindo tempo
- [ ] Versão concorrente com `go` — observar que termina prematuramente
- [ ] Documentar nas anotações: por que falha?
- [ ] Executar `go build -gcflags="-m"` e anotar o que escapa

### Estrutura sugerida:
```
pratica/
├── main.go
├── downloader.go
└── go.mod
```

---

## 📚 Referências da Semana
- [Go By Example — Goroutines](https://gobyexample.com/goroutines)
- [Go Blog — The Go Memory Model](https://go.dev/ref/mem)
- [Go Blog — Go GC: Latency Problem Solved](https://go.dev/blog/ismmkeynote)
- [Go Spec — Goroutines](https://go.dev/ref/spec#Go_statements)
