# 📝 Anotações — Semana 1: A Base Absoluta

> Espaço para anotar descobertas, dúvidas e insights durante o estudo.

---

## 💡 Insights

- **Declaração Explícita:** `var nome string = "Valor"`. Útil quando queremos ser bem claros sobre o tipo ou declarar sem inicializar logo.
- **Escopo do Pacote:** Aprendi que em Go, ficheiros na mesma pasta pertencem ao mesmo pacote e partilham o mesmo espaço de nomes (não se pode ter dois `main` na mesma pasta).

- **Declaração com Inferência de tipo:** `var idade = 22`. Uma das grandes utilidades desse tipo de declaração é a capacidade de utiliza-las fora do nível do escopo de funções.
---

## ❓ Dúvidas

-

---

## ⚠️ Armadilhas / Gotchas

- **Erro `main redeclared`:** Se criar vários ficheiros `.go` na mesma pasta com `package main`, apenas um pode ter a função `main()`. A solução é criar subpastas para cada exercício.

---

## 🔗 Links Úteis Encontrados

<!-- Links adicionais que encontrou durante o estudo -->

-

---

## ✅ Checklist de Aprendizagem

- [ ] Instalei Go e configurei o workspace
- [ ] Entendi `go run` vs `go build`
- [ ] Sei declarar variáveis com `var` e `:=`
- [ ] Conheço os tipos primitivos e seus zero values
- [ ] Entendi `if` com inicialização (`if err := ...`)
- [ ] Domino o `for` em todas as suas formas
- [ ] Sei criar funções com múltiplos retornos
- [ ] Entendi closures
- [ ] Calculadora CLI está funcionando ✨
