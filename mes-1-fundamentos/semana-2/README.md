# Semana 2: Estruturas de Dados e Tipos Customizados

> 🎯 **Objetivo:** Como modelar dados sem classes e herança.

---

## 📖 Teoria

### 1. Arrays vs Slices

#### Arrays (tamanho fixo — raro em Go)
```go
var notas [5]int                    // Array de 5 inteiros (zero value)
cores := [3]string{"r", "g", "b"}  // Inicializado
```

#### Slices (tamanho dinâmico — **use sempre**)
```go
// Criação
s := []int{1, 2, 3}
s := make([]int, 5)      // len=5, cap=5
s := make([]int, 0, 10)  // len=0, cap=10

// Operações fundamentais
len(s)           // Comprimento
cap(s)           // Capacidade (memória alocada)
s = append(s, 4) // Adicionar elemento (pode realocar!)
```

#### ⚠️ Armadilha de Slices (CRUCIAL)
```go
original := []int{1, 2, 3, 4, 5}
fatia := original[1:3]  // [2, 3] — compartilha memória!

fatia[0] = 99           // Altera TAMBÉM o original!
// original agora é [1, 99, 3, 4, 5]

// Solução: usar copy
copia := make([]int, len(fatia))
copy(copia, fatia)
```

### 2. Maps
```go
// Criação
users := make(map[string]int)
users["alice"] = 30
users["bob"] = 25

// Verificar existência
idade, ok := users["alice"]
if !ok {
    fmt.Println("não encontrado")
}

// Deletar
delete(users, "bob")

// Iterar (ordem NÃO garantida!)
for chave, valor := range users {
    fmt.Printf("%s: %d\n", chave, valor)
}
```

### 3. Structs
```go
type Tarefa struct {
    ID        int       `json:"id"`
    Titulo    string    `json:"titulo"`
    Concluida bool      `json:"concluida"`
    CriadaEm  time.Time `json:"criada_em"`
}

// Inicialização
t := Tarefa{
    ID:     1,
    Titulo: "Estudar Go",
}
```

### 4. Ponteiros
```go
x := 42
p := &x      // p aponta para x
fmt.Println(*p) // 42 (desreferência)
*p = 100
fmt.Println(x)  // 100

// Quando usar ponteiro?
// ✅ Quando quer MODIFICAR o valor original
// ✅ Quando o struct é grande (evitar cópia)
// ❌ Quando quer apenas LER um valor pequeno
```

### 5. Métodos (Value vs Pointer Receivers)
```go
// Value receiver — recebe CÓPIA
func (t Tarefa) Resumo() string {
    return fmt.Sprintf("[%d] %s", t.ID, t.Titulo)
}

// Pointer receiver — pode MODIFICAR
func (t *Tarefa) Concluir() {
    t.Concluida = true
}
```

> **Regra prática:** Se pelo menos um método precisa de pointer receiver, use pointer receiver em **todos** os métodos do tipo para manter a consistência.

---

## 🔨 Prática

**Projeto:** Gerenciador de Tarefas (To-Do List) na Memória

```bash
go run main.go add "Estudar Go"
go run main.go list
go run main.go done 1
go run main.go remove 1
```

### Requisitos:
- [ ] Struct `Tarefa` com campos: ID, Titulo, Concluida
- [ ] Struct `GerenciadorTarefas` com um slice de tarefas
- [ ] Método `Adicionar(titulo string)` (pointer receiver)
- [ ] Método `Listar()` que formata e imprime todas
- [ ] Método `Concluir(id int) error` que marca como feita
- [ ] Método `Remover(id int) error` que remove do slice
- [ ] Aceitar comandos via `os.Args`

### Estrutura sugerida:
```
pratica/
├── main.go
├── tarefa.go
└── gerenciador.go
```

---

## 📚 Referências da Semana
- [Tour of Go — Moretypes](https://go.dev/tour/moretypes/1)
- [Go By Example — Slices](https://gobyexample.com/slices)
- [Go By Example — Maps](https://gobyexample.com/maps)
- [Go By Example — Structs](https://gobyexample.com/structs)
- [Go By Example — Pointers](https://gobyexample.com/pointers)
- [Go By Example — Methods](https://gobyexample.com/methods)
- [Go Blog — Slices Internals](https://go.dev/blog/slices-intro)
