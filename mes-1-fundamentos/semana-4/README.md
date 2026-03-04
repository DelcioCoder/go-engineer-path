# Semana 4: Ferramentas Nativas e Testes Básicos

> 🎯 **Objetivo:** Dominar as ferramentas antes de ir para produção.

---

## 📖 Teoria

### 1. Go Modules

```bash
# Inicializar módulo
go mod init github.com/seu-usuario/projeto

# Adicionar dependência
go get github.com/algum/pacote

# Limpar dependências não usadas
go mod tidy

# Verificar checksums
go mod verify
```

Ficheiros gerados:
- `go.mod` — Define o módulo e suas dependências
- `go.sum` — Checksums para garantir integridade

### 2. O Pacote `testing`

```go
// arquivo: calculadora_test.go (DEVE terminar com _test.go)
package calculadora

import "testing"

func TestSoma(t *testing.T) {
    resultado := Somar(2, 3)
    esperado := 5

    if resultado != esperado {
        t.Errorf("Somar(2, 3) = %d; esperado %d", resultado, esperado)
    }
}
```

```bash
go test ./...           # Executar todos os testes
go test -v ./...        # Modo verboso
go test -run TestSoma   # Executar teste específico
go test -cover ./...    # Ver cobertura
go test -coverprofile=coverage.out ./...  # Gerar relatório
go tool cover -html=coverage.out         # Visualizar no browser
```

### 3. Table-Driven Tests (Padrão Idiomático)

```go
func TestDividir(t *testing.T) {
    testes := []struct {
        nome      string
        a, b      float64
        esperado  float64
        esperaErr bool
    }{
        {"divisão normal",     10, 2, 5, false},
        {"divisão decimal",    7, 2, 3.5, false},
        {"divisão por zero",   10, 0, 0, true},
        {"zero dividido",      0, 5, 0, false},
    }

    for _, tt := range testes {
        t.Run(tt.nome, func(t *testing.T) {
            resultado, err := Dividir(tt.a, tt.b)

            if tt.esperaErr && err == nil {
                t.Fatal("esperava erro, mas não recebeu")
            }
            if !tt.esperaErr && err != nil {
                t.Fatalf("não esperava erro, recebeu: %v", err)
            }
            if resultado != tt.esperado {
                t.Errorf("Dividir(%.1f, %.1f) = %.1f; esperado %.1f",
                    tt.a, tt.b, resultado, tt.esperado)
            }
        })
    }
}
```

### 4. Ferramentas de Código

```bash
# Formatação automática (OBRIGATÓRIO — nunca commit sem formatar)
go fmt ./...
# ou
gofmt -w .

# Análise estática (encontra bugs sutis)
go vet ./...

# Exemplos do que go vet detecta:
# - Printf com argumentos errados
# - Variáveis não usadas em loops
# - Métodos de interface mal implementados
```

### 5. Helpers de Teste (t.Helper)

```go
func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper() // Faz o erro apontar para quem chamou, não para esta linha
    if got != want {
        t.Errorf("got %v; want %v", got, want)
    }
}
```

---

## 🔨 Prática

**Projeto:** Escrever Testes para o Gerenciador de Tarefas da Semana 2

### Requisitos:
- [ ] Inicializar módulo Go (`go mod init`)
- [ ] Criar `gerenciador_test.go` com Table-Driven Tests
- [ ] Testar `Adicionar` (com casos válidos)
- [ ] Testar `Concluir` (com ID válido e inválido)
- [ ] Testar `Remover` (com ID válido e inválido)
- [ ] Testar `Listar` usando `bytes.Buffer` como `io.Writer`
- [ ] Executar `go fmt` e `go vet` sem erros
- [ ] Atingir **80%+ de cobertura** (`go test -cover`)

### Estrutura sugerida:
```
pratica/
├── main.go
├── tarefa.go
├── tarefa_test.go
├── gerenciador.go
├── gerenciador_test.go
├── go.mod
└── go.sum
```

---

## 📚 Referências da Semana
- [Go By Example — Testing](https://gobyexample.com/testing-and-benchmarking)
- [Go Blog — Table-Driven Tests](https://go.dev/wiki/TableDrivenTests)
- [Go Blog — Using Go Modules](https://go.dev/blog/using-go-modules)
- [Effective Go — Formatting](https://go.dev/doc/effective_go#formatting)
- [Go Doc — testing package](https://pkg.go.dev/testing)
