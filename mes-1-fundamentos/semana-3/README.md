# Semana 3: O Poder das Interfaces e Tratamento de Erros

> 🎯 **Objetivo:** Dominar o coração da arquitetura e do design de software em Go.

---

## 📖 Teoria

### 1. Tratamento de Erros Explícito

Go **não tem exceções**. Erros são valores retornados explicitamente.

```go
// A interface error
type error interface {
    Error() string
}

// Criando erros
err := errors.New("algo deu errado")
err := fmt.Errorf("usuário %d não encontrado", id)

// Padrão idiomático
resultado, err := fazerAlgo()
if err != nil {
    return fmt.Errorf("falha ao fazer algo: %w", err) // %w = wrap
}
```

#### Erros customizados
```go
type ErroValidacao struct {
    Campo    string
    Mensagem string
}

func (e *ErroValidacao) Error() string {
    return fmt.Sprintf("campo '%s': %s", e.Campo, e.Mensagem)
}

// Verificar tipo do erro
var errVal *ErroValidacao
if errors.As(err, &errVal) {
    fmt.Println("Campo com problema:", errVal.Campo)
}

// Verificar erro específico
if errors.Is(err, ErrNaoEncontrado) {
    // ...
}
```

### 2. defer, panic e recover

```go
// defer — executa no FINAL da função (LIFO)
func lerArquivo(nome string) error {
    f, err := os.Open(nome)
    if err != nil {
        return err
    }
    defer f.Close() // Garantia de que fecha!

    // ... processar arquivo
    return nil
}

// panic — NÃO use para erros normais. Apenas para bugs irrecuperáveis.
// recover — captura um panic (raro em código de aplicação)
```

### 3. Interfaces (O CONCEITO MAIS IMPORTANTE)

```go
// Interfaces são satisfeitas IMPLICITAMENTE (duck typing)
type Escritor interface {
    Escrever(dados []byte) (int, error)
}

// Qualquer tipo com o método Escrever satisfaz Escritor
// NÃO precisa declarar "implements"!
```

#### Interfaces nativas essenciais
```go
// io.Writer — o mais importante
type Writer interface {
    Write(p []byte) (n int, err error)
}
// os.Stdout, *os.File, *bytes.Buffer — todos satisfazem io.Writer!

// io.Reader
type Reader interface {
    Read(p []byte) (n int, err error)
}

// fmt.Stringer — equivale ao toString() de Java
type Stringer interface {
    String() string
}
```

#### O poder: programar contra interfaces
```go
// Aceita QUALQUER coisa que implemente io.Writer
func salvarRelatorio(w io.Writer, dados string) error {
    _, err := fmt.Fprintln(w, dados)
    return err
}

// Agora funciona com arquivo, terminal, buffer...
salvarRelatorio(os.Stdout, "relatório no terminal")
salvarRelatorio(arquivo, "relatório no arquivo")
salvarRelatorio(&buf, "relatório em memória para teste")
```

### 4. Type Assertions e Type Switches

```go
// Type assertion
var i interface{} = "olá"
s, ok := i.(string) // ok = true, s = "olá"
n, ok := i.(int)    // ok = false, n = 0

// Type switch
func descrever(i interface{}) string {
    switch v := i.(type) {
    case string:
        return "string: " + v
    case int:
        return fmt.Sprintf("inteiro: %d", v)
    case error:
        return "erro: " + v.Error()
    default:
        return "tipo desconhecido"
    }
}
```

---

## 🔨 Prática

**Projeto:** Estender a Calculadora (Semana 1) ou o Gerenciador (Semana 2) com `io.Writer`

### Opção A: Calculadora com io.Writer
```go
func calcular(w io.Writer, a, b float64, op string) error {
    // Escreve resultado em qualquer Writer
}

// Terminal
calcular(os.Stdout, 10, 5, "+")

// Arquivo
f, _ := os.Create("resultado.txt")
calcular(f, 10, 5, "+")
```

### Opção B: Gerenciador de Tarefas com io.Writer
```go
func (g *GerenciadorTarefas) Listar(w io.Writer) error {
    // Lista tarefas escrevendo em qualquer Writer
}
```

### Requisitos:
- [ ] Refatorar para aceitar `io.Writer` como parâmetro
- [ ] Funcionar com `os.Stdout` (terminal)
- [ ] Funcionar com `*os.File` (arquivo)
- [ ] Funcionar com `*bytes.Buffer` (memória/testes)
- [ ] Implementar a interface `fmt.Stringer` no struct `Tarefa`
- [ ] Usar `errors.New` e `fmt.Errorf` com `%w` para wrap de erros
- [ ] Criar pelo menos um erro customizado com struct

---

## 📚 Referências da Semana
- [Tour of Go — Methods and Interfaces](https://go.dev/tour/methods/1)
- [Go By Example — Interfaces](https://gobyexample.com/interfaces)
- [Go By Example — Errors](https://gobyexample.com/errors)
- [Go Blog — Error Handling](https://go.dev/blog/error-handling-and-go)
- [Go Blog — Errors are Values](https://go.dev/blog/errors-are-values)
- [Effective Go — Interfaces](https://go.dev/doc/effective_go#interfaces)
