# Semana 10: Persistência de Dados e Clean Architecture

> 🎯 **Objetivo:** Sair da memória RAM e organizar o projeto profissionalmente.

---

## 📖 Teoria

### 1. database/sql — O Driver Oficial

```go
import (
    "database/sql"
    _ "github.com/jackc/pgx/v5/stdlib"  // Driver PostgreSQL
)

func conectar() (*sql.DB, error) {
    db, err := sql.Open("pgx", "postgres://user:pass@localhost:5432/todos?sslmode=disable")
    if err != nil {
        return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
    }

    // Testar conexão
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("erro ao conectar ao banco: %w", err)
    }

    // Configurar pool de conexões
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    return db, nil
}
```

### 2. Operações CRUD com database/sql

```go
// INSERT
func (r *TarefaRepo) Criar(ctx context.Context, t *Tarefa) error {
    query := `INSERT INTO tarefas (titulo, concluida) VALUES ($1, $2) RETURNING id`
    return r.db.QueryRowContext(ctx, query, t.Titulo, t.Concluida).Scan(&t.ID)
}

// SELECT (múltiplos)
func (r *TarefaRepo) Listar(ctx context.Context) ([]Tarefa, error) {
    rows, err := r.db.QueryContext(ctx, `SELECT id, titulo, concluida FROM tarefas`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tarefas []Tarefa
    for rows.Next() {
        var t Tarefa
        if err := rows.Scan(&t.ID, &t.Titulo, &t.Concluida); err != nil {
            return nil, err
        }
        tarefas = append(tarefas, t)
    }
    return tarefas, rows.Err() // IMPORTANTE: verificar rows.Err()
}

// SELECT (único)
func (r *TarefaRepo) Buscar(ctx context.Context, id int) (*Tarefa, error) {
    var t Tarefa
    err := r.db.QueryRowContext(ctx,
        `SELECT id, titulo, concluida FROM tarefas WHERE id = $1`, id,
    ).Scan(&t.ID, &t.Titulo, &t.Concluida)

    if errors.Is(err, sql.ErrNoRows) {
        return nil, ErrNaoEncontrado
    }
    return &t, err
}
```

### 3. Transações SQL

```go
func (r *TarefaRepo) TransferirProjeto(ctx context.Context, tarefaID, projetoID int) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback() // Rollback se não fizer Commit

    _, err = tx.ExecContext(ctx, `UPDATE tarefas SET projeto_id = $1 WHERE id = $2`, projetoID, tarefaID)
    if err != nil {
        return err
    }

    _, err = tx.ExecContext(ctx, `UPDATE projetos SET total_tarefas = total_tarefas + 1 WHERE id = $1`, projetoID)
    if err != nil {
        return err
    }

    return tx.Commit() // Só comita se tudo deu certo
}
```

### 4. Estruturação de Projetos Go

```
projeto/
├── cmd/
│   └── api/
│       └── main.go           ← Ponto de entrada
├── internal/
│   ├── handler/
│   │   └── tarefa_handler.go ← Handlers HTTP
│   ├── service/
│   │   └── tarefa_service.go ← Lógica de negócio
│   ├── repository/
│   │   └── tarefa_repo.go    ← Acesso a dados
│   └── model/
│       └── tarefa.go         ← Structs/Modelos
├── go.mod
├── go.sum
└── README.md
```

> **`internal/`** — Código que NÃO pode ser importado por outros módulos!

### 5. Padrão Repository + Service

```go
// Interface do Repositório (CONTRATO)
type TarefaRepository interface {
    Criar(ctx context.Context, t *Tarefa) error
    Listar(ctx context.Context) ([]Tarefa, error)
    Buscar(ctx context.Context, id int) (*Tarefa, error)
    Atualizar(ctx context.Context, t *Tarefa) error
    Deletar(ctx context.Context, id int) error
}

// Implementação Postgres
type postgresRepo struct {
    db *sql.DB
}

func NovoPostgresRepo(db *sql.DB) TarefaRepository {
    return &postgresRepo{db: db}
}

// Implementação Mock (para testes!)
type mockRepo struct {
    tarefas map[int]*Tarefa
}

func NovoMockRepo() TarefaRepository {
    return &mockRepo{tarefas: make(map[int]*Tarefa)}
}

// Service (lógica de negócio — NÃO sabe do banco!)
type TarefaService struct {
    repo TarefaRepository  // Depende da INTERFACE, não da implementação
}

func NovoTarefaService(repo TarefaRepository) *TarefaService {
    return &TarefaService{repo: repo}
}
```

---

## 🔨 Prática

**Projeto:** Migrar a API REST da Semana 9 para PostgreSQL com Clean Architecture

### Requisitos:
- [ ] Estruturar pastas: `cmd/`, `internal/handler/`, `internal/service/`, `internal/repository/`
- [ ] Criar interface `TarefaRepository`
- [ ] Implementar `postgresRepo` com `database/sql`
- [ ] Implementar `mockRepo` para testes
- [ ] Criar `TarefaService` que depende da interface
- [ ] Handlers chamam Service, Service chama Repository
- [ ] Usar `context.Context` em todas as queries
- [ ] Configurar pool de conexões
- [ ] SQL: `CREATE TABLE tarefas (id SERIAL PRIMARY KEY, titulo TEXT NOT NULL, concluida BOOLEAN DEFAULT FALSE)`

---

## 📚 Referências da Semana
- [Go Doc — database/sql](https://pkg.go.dev/database/sql)
- [Go Wiki — SQLInterface](https://go.dev/wiki/SQLInterface)
- [Go Blog — Organizing Go Code](https://go.dev/blog/organizing-go-code)
- [Go Standard Project Layout](https://github.com/golang-standards/project-layout)
- [pgx — PostgreSQL Driver](https://github.com/jackc/pgx)
