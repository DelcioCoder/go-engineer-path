# Semana 1: A Base Absoluta e Sintaxe

> 🎯 **Objetivo:** Entender as regras do jogo e por que Go é tão minimalista.

---

## 📖 Teoria

### 1. Instalação e Workspace
- Instalar Go via [go.dev/dl](https://go.dev/dl/)
- Configurar `GOPATH` e `GOROOT` (Go modules eliminam a necessidade de `GOPATH` clássico)
- Verificar: `go version`
- Criar o primeiro módulo: `go mod init github.com/seu-usuario/calculadora`

### 2. Compilação e Execução
```bash
go run main.go      # Compila e executa (dev)
go build -o app     # Compila binário nativo
./app               # Executa o binário
```

### 3. Declaração de Variáveis
```go
// Declaração explícita
var nome string = "Dercio"
var idade int = 25

// Declaração curta (apenas dentro de funções)
cidade := "Luanda"
ativo := true

// Zero values (valores padrão)
var x int       // 0
var s string    // ""
var b bool      // false
```

### 4. Tipos Primitivos
| Tipo | Exemplos | Zero Value |
|------|----------|------------|
| `int`, `int8`, `int16`, `int32`, `int64` | `42` | `0` |
| `uint`, `uint8` (byte), `uint16`, `uint32`, `uint64` | `42` | `0` |
| `float32`, `float64` | `3.14` | `0.0` |
| `string` | `"olá"` | `""` |
| `bool` | `true`, `false` | `false` |
| `rune` (alias para `int32`) | `'A'` | `0` |

### 5. Estruturas de Controle

#### if/else com inicialização
```go
// Padrão idiomático Go — variável no escopo do if
if err := fazerAlgo(); err != nil {
    log.Fatal(err)
}

// if/else clássico
if x > 10 {
    fmt.Println("grande")
} else if x > 5 {
    fmt.Println("médio")
} else {
    fmt.Println("pequeno")
}
```

#### switch (sem break — cai automaticamente!)
```go
switch op {
case "+":
    resultado = a + b
case "-":
    resultado = a - b
default:
    fmt.Println("operação desconhecida")
}
```

### 6. O Único Loop: `for`
```go
// Clássico
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Estilo "while"
for condição {
    // ...
}

// Infinito
for {
    // break para sair
}

// Range (iteração sobre coleções)
for i, v := range slice {
    fmt.Printf("índice: %d, valor: %v\n", i, v)
}
```

### 7. Funções
```go
// Múltiplos retornos
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}

// Função anônima
somar := func(a, b int) int {
    return a + b
}

// Closure (captura variável externa)
func contador() func() int {
    n := 0
    return func() int {
        n++
        return n
    }
}
```

---

## 🔨 Prática

**Projeto:** Calculadora via CLI

Criar uma calculadora que aceita argumentos via `os.Args`:

```bash
go run main.go 10 + 5
# Saída: 10 + 5 = 15

go run main.go 20 / 0
# Saída: erro: divisão por zero
```

### Requisitos:
- [ ] Aceitar 3 argumentos: `número1 operador número2`
- [ ] Suportar operações: `+`, `-`, `*`, `/`
- [ ] Tratar divisão por zero com mensagem de erro
- [ ] Converter strings para números (`strconv.ParseFloat`)
- [ ] Usar múltiplos retornos para resultado + erro
- [ ] Validar número de argumentos

### Estrutura sugerida:
```
pratica/
├── main.go
```

---

## 📚 Referências da Semana
- [Tour of Go — Basics](https://go.dev/tour/basics/1)
- [Tour of Go — Flow Control](https://go.dev/tour/flowcontrol/1)
- [Go By Example — Variables](https://gobyexample.com/variables)
- [Go By Example — For](https://gobyexample.com/for)
- [Go By Example — Functions](https://gobyexample.com/functions)
- [Go By Example — Closures](https://gobyexample.com/closures)
