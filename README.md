# 🗺️ O Caminho do Engenheiro Go

> **Plano de 12 semanas** para dominar Go de forma idiomática e produtiva — do zero ao deploy em produção.

**Início:** Março 2026  
**Previsão de término:** Junho 2026

---

## 📊 Progresso Geral

### 🟢 Mês 1: Fundamentos e Pensamento Idiomático
| Semana | Tema | Status |
|--------|------|--------|
| 1 | A Base Absoluta e Sintaxe | ⬜ |
| 2 | Estruturas de Dados e Tipos Customizados | ⬜ |
| 3 | O Poder das Interfaces e Tratamento de Erros | ⬜ |
| 4 | Ferramentas Nativas e Testes Básicos | ⬜ |

### 🟡 Mês 2: Debaixo dos Panos e Concorrência Real
| Semana | Tema | Status |
|--------|------|--------|
| 5 | Memória e Fundamentos de Concorrência | ⬜ |
| 6 | Sincronização e Canais | ⬜ |
| 7 | Concorrência Avançada e Contextos | ⬜ |
| 8 | Profiling e Generics | ⬜ |

### 🔴 Mês 3: Arquitetura, Backend e Sistemas
| Semana | Tema | Status |
|--------|------|--------|
| 9 | Servidores HTTP e Redes (Standard Library) | ⬜ |
| 10 | Persistência de Dados e Clean Architecture | ⬜ |
| 11 | Operações de Produção | ⬜ |
| 12 | Projeto Final de Portfólio | ⬜ |

> **Legenda:** ⬜ Não iniciado · 🔄 Em progresso · ✅ Concluído

---

## 🗂️ Estrutura do Repositório

```
go-engineer-path/
├── README.md                         ← Você está aqui
├── recursos.md                       ← Links oficiais organizados
├── mes-1-fundamentos/
│   ├── semana-1/                     ← Sintaxe e Calculadora CLI
│   │   ├── README.md
│   │   ├── anotacoes.md
│   │   └── pratica/
│   ├── semana-2/                     ← Structs, Slices e To-Do List
│   │   ├── README.md
│   │   ├── anotacoes.md
│   │   └── pratica/
│   ├── semana-3/                     ← Interfaces e Erros
│   │   ├── README.md
│   │   ├── anotacoes.md
│   │   └── pratica/
│   └── semana-4/                     ← Testes e Ferramentas
│       ├── README.md
│       ├── anotacoes.md
│       └── pratica/
├── mes-2-concorrencia/
│   ├── semana-5/                     ← Goroutines e Memória
│   ├── semana-6/                     ← Channels e sync
│   ├── semana-7/                     ← context e select
│   └── semana-8/                     ← Profiling e Generics
├── mes-3-arquitetura/
│   ├── semana-9/                     ← net/http puro
│   ├── semana-10/                    ← database/sql e Clean Arch
│   ├── semana-11/                    ← Graceful Shutdown e slog
│   └── semana-12/                    ← Projeto Final
└── projeto-final/                    ← Event Processor completo
    ├── cmd/
    ├── internal/
    ├── Dockerfile
    ├── docker-compose.yml
    └── README.md
```

---

## 🎯 Filosofia

1. **Sem frameworks no início.** Domine a standard library antes de qualquer abstração.
2. **Código a cada semana.** Teoria sem prática é ilusão.
3. **Testes desde o dia 1.** Um engenheiro Go escreve testes — sempre.
4. **Simplicidade > Complexidade.** Go foi feito para ser simples. Respeite isso.

---

## 📚 Recursos Essenciais

| Recurso | Tipo | Link |
|---------|------|------|
| Tour of Go | Interativo | [go.dev/tour](https://go.dev/tour) |
| Effective Go | Guia Oficial | [go.dev/doc/effective_go](https://go.dev/doc/effective_go) |
| Go By Example | Exemplos | [gobyexample.com](https://gobyexample.com/) |
| 100 Go Mistakes | Livro | [100go.co](https://100go.co/) |
| Go Playground | Sandbox | [go.dev/play](https://go.dev/play/) |
| Go Blog | Artigos | [go.dev/blog](https://go.dev/blog/) |

> Veja o ficheiro [`recursos.md`](recursos.md) para uma lista completa e organizada por tema.