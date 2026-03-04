# 📚 Kit de Sobrevivência — Recursos Oficiais

> Todos os links que um engenheiro Go precisa, organizados por categoria.

---

## 🏫 Aprendizagem Fundamental

| Recurso | Descrição |
|---------|-----------|
| [Tour of Go](https://go.dev/tour) | Tutorial interativo oficial — **comece aqui** |
| [Effective Go](https://go.dev/doc/effective_go) | Guia de escrita idiomática — **leia e releia** |
| [Go By Example](https://gobyexample.com/) | Padrões práticos com exemplos concisos |
| [Go Playground](https://go.dev/play/) | Sandbox online para testar código |

---

## 📖 Livros Recomendados

| Livro | Autor | Por quê? |
|-------|-------|----------|
| [100 Go Mistakes and How to Avoid Them](https://100go.co/) | Teiva Harsanyi | Transforma juniores em seniores — **essencial** |
| [The Go Programming Language](https://www.gopl.io/) | Donovan & Kernighan | A "bíblia" do Go |
| [Learning Go](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/) | Jon Bodner | Excelente para quem vem de outras linguagens |
| [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) | Katherine Cox-Buday | Profundidade em concorrência |

---

## 📝 Documentação Oficial

| Recurso | Link |
|---------|------|
| Especificação da Linguagem | [go.dev/ref/spec](https://go.dev/ref/spec) |
| Standard Library | [pkg.go.dev/std](https://pkg.go.dev/std) |
| Go Blog | [go.dev/blog](https://go.dev/blog/) |
| Go Wiki | [go.dev/wiki](https://go.dev/wiki/) |
| Release Notes | [go.dev/doc/devel/release](https://go.dev/doc/devel/release) |

---

## 🔧 Ferramentas Essenciais

| Ferramenta | Comando | Finalidade |
|------------|---------|-----------|
| `go fmt` | `go fmt ./...` | Formatação automática do código |
| `go vet` | `go vet ./...` | Análise estática de erros comuns |
| `go test` | `go test ./...` | Executar testes unitários |
| `go test -cover` | `go test -cover ./...` | Cobertura de testes |
| `go test -race` | `go test -race ./...` | Detector de data races |
| `go test -bench` | `go test -bench=. ./...` | Benchmarks |
| `go build -gcflags="-m"` | — | Escape Analysis |
| `go tool pprof` | — | Profiling de CPU/Memória |
| `golangci-lint` | `golangci-lint run` | Linter completo (instalar separadamente) |

---

## 🧪 Pacotes Nativos Importantes (por semana)

### Mês 1
| Pacote | Semana | Uso |
|--------|--------|-----|
| `fmt` | 1 | Entrada/Saída formatada |
| `os` | 1 | Argumentos CLI (`os.Args`) |
| `strconv` | 1 | Conversão de string para números |
| `errors` | 3 | Criação de erros customizados |
| `io` | 3 | Interfaces `Reader` e `Writer` |
| `testing` | 4 | Testes unitários |

### Mês 2
| Pacote | Semana | Uso |
|--------|--------|-----|
| `sync` | 5-6 | `WaitGroup`, `Mutex`, `RWMutex` |
| `context` | 7 | Cancelamento e timeouts |
| `time` | 5-7 | Medição de tempo e `time.After` |
| `runtime` | 5 | Informações do runtime/scheduler |
| `slices` | 8 | Funções genéricas para slices |

### Mês 3
| Pacote | Semana | Uso |
|--------|--------|-----|
| `net/http` | 9 | Servidor HTTP e roteamento |
| `encoding/json` | 9 | Serialização JSON |
| `database/sql` | 10 | Acesso a banco de dados |
| `os/signal` | 11 | Captura de sinais do SO |
| `log/slog` | 11 | Logging estruturado |

---

## 🌐 Comunidades

| Comunidade | Link |
|------------|------|
| Gophers Slack | [gophers.slack.com](https://gophers.slack.com) |
| r/golang | [reddit.com/r/golang](https://www.reddit.com/r/golang/) |
| Go Forum | [forum.golangbridge.org](https://forum.golangbridge.org/) |
| Go Weekly Newsletter | [golangweekly.com](https://golangweekly.com/) |

---

## 🎥 Canais e Talks

| Canal/Talk | Link |
|------------|------|
| GopherCon (YouTube) | [youtube.com/@GopherAcademy](https://www.youtube.com/@GopherAcademy) |
| justforfunc | [youtube.com/@justforfunc](https://www.youtube.com/@justforfunc) |
| Rob Pike - "Simplicity is Complicated" | [youtube.com/watch?v=rFejpH_tAHM](https://www.youtube.com/watch?v=rFejpH_tAHM) |
| Rob Pike - "Concurrency is not Parallelism" | [youtube.com/watch?v=oV9rvDllKEg](https://www.youtube.com/watch?v=oV9rvDllKEg) |
