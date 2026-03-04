# Semana 9: Servidores HTTP e Redes (Standard Library)

> 🎯 **Objetivo:** Não toque em frameworks (Gin, Fiber, etc) ainda. Domine a standard library.

---

## 📖 Teoria

### 1. Criando um Servidor HTTP

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Olá, Go!")
    })

    log.Println("Servidor rodando em :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

### 2. Roteamento com ServeMux (Go 1.22+)

```go
mux := http.NewServeMux()

// Métodos explícitos + path variables (Go 1.22+)
mux.HandleFunc("GET /todos", listarTodos)
mux.HandleFunc("POST /todos", criarTodo)
mux.HandleFunc("GET /todos/{id}", buscarTodo)
mux.HandleFunc("PUT /todos/{id}", atualizarTodo)
mux.HandleFunc("DELETE /todos/{id}", deletarTodo)

// Acessar path variable
func buscarTodo(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")  // Go 1.22+
    // ...
}
```

### 3. Parsing de JSON (encoding/json)

```go
// Struct → JSON (Marshal/Encode)
type Tarefa struct {
    ID        int    `json:"id"`
    Titulo    string `json:"titulo"`
    Concluida bool   `json:"concluida,omitempty"`
}

// Encoder (escreve direto no Writer — mais eficiente)
func responderJSON(w http.ResponseWriter, status int, dados interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(dados)
}

// Decoder (lê direto do Reader — mais eficiente e seguro)
func lerJSON(r *http.Request, destino interface{}) error {
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields() // Rejeita campos extras!
    return decoder.Decode(destino)
}
```

### 4. Middlewares Puros

```go
// Middleware é uma função que envolve um http.Handler
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        inicio := time.Now()

        // Chama o próximo handler
        next.ServeHTTP(w, r)

        duracao := time.Since(inicio)
        log.Printf("%s %s %v", r.Method, r.URL.Path, duracao)
    })
}

// Compor middlewares
handler := loggingMiddleware(corsMiddleware(mux))
http.ListenAndServe(":8080", handler)
```

### 5. A Interface http.Handler

```go
// Toda a magia do net/http vem desta interface simples
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

// http.HandlerFunc é um adaptador que transforma uma função em Handler
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

---

## 🔨 Prática

**Projeto:** API REST para To-Do List (Standard Library pura)

### Endpoints:
| Método | Rota | Descrição |
|--------|------|-----------|
| `GET` | `/todos` | Listar todas as tarefas |
| `POST` | `/todos` | Criar nova tarefa |
| `GET` | `/todos/{id}` | Buscar tarefa por ID |
| `PUT` | `/todos/{id}` | Atualizar tarefa |
| `DELETE` | `/todos/{id}` | Remover tarefa |

### Requisitos:
- [ ] Servidor HTTP com `net/http` e `ServeMux`
- [ ] Roteamento com métodos e path variables (Go 1.22+)
- [ ] Request/Response em JSON com `encoding/json`
- [ ] Usar `json.Decoder` (não `json.Unmarshal`)
- [ ] Middleware de logging: método, rota, duração em ms
- [ ] Retornar status codes corretos (200, 201, 404, 400, 500)
- [ ] Armazenamento na memória (por enquanto)

### Estrutura sugerida:
```
pratica/
├── main.go          (setup do servidor)
├── handler.go       (handlers HTTP)
├── middleware.go     (logging middleware)
├── modelo.go        (structs)
├── storage.go       (armazenamento em memória)
└── go.mod
```

---

## 📚 Referências da Semana
- [Go Doc — net/http](https://pkg.go.dev/net/http)
- [Go Doc — encoding/json](https://pkg.go.dev/encoding/json)
- [Go Blog — Routing Enhancements (1.22)](https://go.dev/blog/routing-enhancements)
- [Go By Example — HTTP Server](https://gobyexample.com/http-server)
- [Go By Example — JSON](https://gobyexample.com/json)
- [Go Wiki — HTTP Handler Patterns](https://go.dev/wiki/LearnServerProgramming)
