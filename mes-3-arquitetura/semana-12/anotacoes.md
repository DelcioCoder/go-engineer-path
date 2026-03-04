# 📝 Anotações — Semana 12: Projeto Final

> Espaço para anotar decisões de design, gargalos encontrados e lições aprendidas.

---

## 💡 Decisões de Design

-

---

## 🏗️ Desafios de Arquitetura

-

---

## 🐛 Bugs Encontrados e Resolvidos

-

---

## 📊 Métricas de Performance

<!-- Resultados de benchmarks, profiling, etc. -->

-

---

## ✅ Checklist Final

### API
- [ ] POST /events funcionando
- [ ] GET /events listando do banco
- [ ] GET /health retornando status

### Concorrência
- [ ] Worker pool estável sob carga
- [ ] Context timeout protegendo workers
- [ ] Sem data races (`go test -race`)

### Infraestrutura
- [ ] Docker build funcionando
- [ ] docker-compose up funcional
- [ ] Imagem final < 20MB (multistage)

### Qualidade
- [ ] Testes com 80%+ cobertura
- [ ] Zero `go vet` warnings
- [ ] Código formatado com `go fmt`
- [ ] README espetacular ✨
