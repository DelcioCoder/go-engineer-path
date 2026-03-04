# Semana 8: Profiling e Generics

> 🎯 **Objetivo:** Monitorar a saúde do software e usar tipos paramétricos.

---

## 📖 Teoria

### 1. Benchmarks Nativos

```go
// arquivo: algo_test.go
func BenchmarkProcessar(b *testing.B) {
    dados := prepararDados()

    b.ResetTimer() // Ignora tempo de setup
    for i := 0; i < b.N; i++ {
        processar(dados)
    }
}
```

```bash
# Executar benchmarks
go test -bench=. ./...

# Com alocações de memória
go test -bench=. -benchmem ./...

# Saída:
# BenchmarkProcessar-8  1000000  1234 ns/op  256 B/op  3 allocs/op
#                                 ^^^^^^^^^^  ^^^^^^^^  ^^^^^^^^^^^
#                                 tempo/op    bytes/op  alocações/op
```

### 2. Profiling com pprof

```bash
# Gerar CPU profile
go test -bench=. -cpuprofile=cpu.out ./...

# Gerar Memory profile
go test -bench=. -memprofile=mem.out ./...

# Analisar com pprof interativo
go tool pprof cpu.out
# (pprof) top10          → Top 10 funções
# (pprof) list funcNome  → Código anotado
# (pprof) web            → Flamegraph no browser

# HTTP profiling (para servidores em execução)
import _ "net/http/pprof"
// Aceder: http://localhost:6060/debug/pprof/
```

### 3. Flamegraphs

Um flamegraph mostra **visualmente** onde o programa gasta tempo:
- **Eixo X:** proporção do tempo total
- **Eixo Y:** profundidade da call stack
- **Funções largas:** consomem mais tempo → investigar!

### 4. Go Generics (Go 1.18+)

```go
// Antes de generics — duplicação
func SomarInts(nums []int) int { ... }
func SomarFloats(nums []float64) float64 { ... }

// Com generics — uma função serve para tudo
func Somar[T int | float64](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

// Usando constraints do pacote cmp/constraints
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

### 5. Funções Genéricas Nativas (slices, maps)

```go
import "slices"

// Ordenar com função customizada
pessoas := []Pessoa{{Nome: "Ana", Idade: 30}, {Nome: "Bob", Idade: 25}}
slices.SortFunc(pessoas, func(a, b Pessoa) int {
    return a.Idade - b.Idade
})

// Outras funções úteis
slices.Contains(s, valor)
slices.Index(s, valor)
slices.Reverse(s)
```

---

## 🔨 Prática

### Parte 1: Encontrar Gargalos com Profiling
- [ ] Adicionar benchmarks ao código das semanas anteriores
- [ ] Gerar CPU profile e identificar funções lentas
- [ ] Gerar Memory profile e verificar alocações
- [ ] Visualizar flamegraph com `go tool pprof`
- [ ] Documentar gargalos encontrados

### Parte 2: Generics
- [ ] Criar função genérica `Filtrar[T any](slice []T, fn func(T) bool) []T`
- [ ] Usar `slices.SortFunc` para ordenar structs
- [ ] Comparar performance genérica vs concreta com benchmarks

---

## 📚 Referências da Semana
- [Go By Example — Generics](https://gobyexample.com/generics)
- [Go Blog — An Introduction to Generics](https://go.dev/blog/intro-generics)
- [Go Blog — Profiling Go Programs](https://go.dev/blog/pprof)
- [Go Doc — testing (Benchmarks)](https://pkg.go.dev/testing#hdr-Benchmarks)
- [Go Doc — slices package](https://pkg.go.dev/slices)
- [Go Blog — Diagnostics](https://go.dev/doc/diagnostics)
