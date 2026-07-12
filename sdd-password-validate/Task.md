# Tarefas - Gerador de Senhas com Validação

## Fase 1: Setup do Projeto

### Tarefa 1.1: Criar Estrutura de Pastas
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 0.5 dia

**Descrição:**
Criar a estrutura de pastas conforme Clean Architecture.

**Subtarefas:**
- [ ] Criar pasta `cmd/api/`
- [ ] Criar pasta `internal/domain/`
- [ ] Criar pasta `internal/usecase/`
- [ ] Criar pasta `internal/interfaces/`
- [ ] Criar pasta `internal/infrastructure/handlers/`
- [ ] Criar pasta `internal/infrastructure/services/`
- [ ] Criar pasta `internal/infrastructure/config/`
- [ ] Criar pasta `internal/infrastructure/router/`
- [ ] Criar pasta `internal/pkg/errors/`

**Critérios de Aceite:**
- Estrutura de pastas criada conforme especificação
- Pastas vazias prontas para receber código

---

### Tarefa 1.2: Inicializar Módulo Go
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 0.25 dia

**Descrição:**
Inicializar módulo Go e configurar dependências básicas.

**Subtarefas:**
- [ ] Executar `go mod init password-generator`
- [ ] Criar `.gitignore` adequado para Go
- [ ] Verificar versão do Go instalada

**Critérios de Aceite:**
- `go.mod` criado corretamente
- `.gitignore` configurado

---

### Tarefa 1.3: Adicionar Dependências
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 0.25 dia

**Descrição:**
Adicionar todas as dependências necessárias ao projeto.

**Subtarefas:**
- [ ] Executar `go get github.com/gin-gonic/gin`
- [ ] Executar `go get github.com/spf13/viper`
- [ ] Executar `go get github.com/stretchr/testify`
- [ ] Executar `go get github.com/swaggo/swag`
- [ ] Executar `go get github.com/swaggo/gin-swagger`
- [ ] Executar `go get golang.org/x/crypto`
- [ ] Executar `go mod tidy`

**Critérios de Aceite:**
- `go.mod` com todas as dependências
- `go.sum` atualizado

---

### Tarefa 1.4: Criar Makefile
**Status:** Pendente
**Prioridade:** Média
**Estimativa:** 0.25 dia

**Descrição:**
Criar Makefile com comandos úteis para desenvolvimento.

**Subtarefas:**
- [ ] Criar targets: `build`, `run`, `test`, `lint`
- [ ] Criar target para Docker
- [ ] Criar target para Swagger

**Critérios de Aceite:**
- `make build` funciona
- `make test` funciona
- `make run` funciona

---

## Fase 2: Camada Domain

### Tarefa 2.1: Criar Entidade Password
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 0.5 dia

**Descrição:**
Implementar a entidade Password com suas propriedades e métodos.

**Arquivo:** `internal/domain/password.go`

**Subtarefas:**
- [ ] Definir struct Password
- [ ] Implementar método NewPassword
- [ ] Implementar método String (mascarado para logs)
- [ ] Implementar validações básicas

**Critérios de Aceite:**
- Struct Password implementada
- Método de mascaramento funcional

---

### Tarefa 2.2: Criar Entidade ValidationResult
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 0.5 dia

**Descrição:**
Implementar a entidade ValidationResult para retorno de validações.

**Arquivo:** `internal/domain/validation.go`

**Subtarefas:**
- [ ] Definir struct ValidationResult
- [ ] Implementar método NewValidationResult
- [ ] Implementar métodos de classificação de força
- [ ] Implementar cálculo de entropia

**Critérios de Aceite:**
- Struct ValidationResult implementada
- Cálculo de entropia correto
- Classificação de força adequada

---

## Fase 3: Casos de Uso

### Tarefa 3.1: Implementar GeneratePassword
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 1 dia

**Descrição:**
Implementar o caso de uso para geração de senhas.

**Arquivo:** `internal/usecase/generate_password.go`

**Subtarefas:**
- [ ] Definir struct GeneratePasswordUseCase
- [ ] Implementar método Execute
- [ ] Integrar com interface PasswordGenerator
- [ ] Validar configuração de entrada
- [ ] Retornar senha + validação

**Critérios de Aceite:**
- Geração de senha funcional
- Integração com interface correta
- Validação de entrada implementada

---

### Tarefa 3.2: Implementar ValidatePassword
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 1 dia

**Descrição:**
Implementar o caso de uso para validação de senhas.

**Arquivo:** `internal/usecase/validate_password.go`

**Subtarefas:**
- [ ] Definir struct ValidatePasswordUseCase
- [ ] Implementar método Execute
- [ ] Integrar com interface PasswordValidator
- [ ] Implementar regras de validação:
  - Comprimento mínimo/máximo
  - Caracteres especiais
  - Repetição consecutiva
- [ ] Calcular entropia
- [ ] Classificar força

**Critérios de Aceite:**
- Validação de senha funcional
- Todas as regras implementadas
- Entropia calculada corretamente

---

## Fase 4: Serviços de Infraestrutura

### Tarefa 4.1: Implementar CryptoGenerator
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 1 dia

**Descrição:**
Implementar o gerador de senhas usando crypto/rand.

**Arquivo:** `internal/infrastructure/services/crypto_generator.go`

**Subtarefas:**
- [ ] Implementar interface PasswordGenerator
- [ ] Usar crypto/rand para aleatoriedade
- [ ] Suportar maiúsculas, minúsculas, números, símbolos
- [ ] Implementar configuração de caracteres

**Critérios de Aceite:**
- Geração usando crypto/rand
- Suporte a todos os tipos de caracteres
- Aleatoriedade verificável

---

### Tarefa 4.2: Implementar RuleValidator
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 1 dia

**Descrição:**
Implementar o validador de regras de senhas.

**Arquivo:** `internal/infrastructure/services/rule_validator.go`

**Subtarefas:**
- [ ] Implementar interface PasswordValidator
- [ ] Implementar validação de comprimento
- [ ] Implementar validação de caracteres especiais
- [ ] Implementar validação de repetição consecutiva
- [ ] Calcular entropia
- [ ] Classificar força

**Critérios de Aceite:**
- Todas as regras implementadas
- Cálculo de entropia correto
- Classificação adequada

---

## Fase 5: Handlers e Router

### Tarefa 5.1: Implementar PasswordHandler
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 0.5 dia

**Descrição:**
Implementar handlers HTTP para a API.

**Arquivo:** `internal/infrastructure/handlers/password_handler.go`

**Subtarefas:**
- [ ] Implementar handler GeneratePassword
- [ ] Implementar handler ValidatePassword
- [ ] Implementar handler HealthCheck
- [ ] Tratar erros e respostas HTTP
- [ ] Validar bodies de requisição

**Critérios de Aceite:**
- Handlers funcionando
- Respostas HTTP corretas
- Tratamento de erros implementado

---

### Tarefa 5.2: Configurar Router
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 0.5 dia

**Descrição:**
Configurar rotas da API.

**Arquivo:** `internal/infrastructure/router/router.go`

**Subtarefas:**
- [ ] Configurar rotas POST `/api/v1/password/generate`
- [ ] Configurar rotas POST `/api/v1/password/validate`
- [ ] Configurar rota GET `/health`
- [ ] Configurar middleware de CORS

**Critérios de Aceite:**
- Rotas configuradas corretamente
- CORS funcional

---

### Tarefa 5.3: Criar Entry Point
**Status:** Pendente
**Prioridade:** Alta
**Estimativa:** 0.25 dia

**Descrição:**
Criar arquivo principal de inicialização.

**Arquivo:** `cmd/api/main.go`

**Subtarefas:**
- [ ] Carregar configuração
- [ ] Inicializar dependências
- [ ] Configurar router
- [ ] Iniciar servidor

**Critérios de Aceite:**
- Aplicação inicia corretamente
- Servidor escuta na porta configurada

---

## Fase 6: Configuração e Docker

### Tarefa 6.1: Implementar Configuração
**Status:** Pendente
**Prioridade:** Média
**Estimativa:** 0.5 dia

**Descrição:**
Implementar sistema de configuração via variáveis de ambiente.

**Arquivo:** `internal/infrastructure/config/config.go`

**Subtarefas:**
- [ ] Configurar Viper para variáveis de ambiente
- [ ] Definir valores padrão
- [ ] Criar structs de configuração
- [ ] Validar configuração

**Critérios de Aceite:**
- Configuração via env vars funcionando
- Valores padrão definidos

---

### Tarefa 6.2: Criar Dockerfile
**Status:** Pendente
**Prioridade:** Média
**Estimativa:** 0.25 dia

**Descrição:**
Criar Dockerfile multi-stage para a aplicação.

**Arquivo:** `Dockerfile`

**Subtarefas:**
- [ ] Criar stage de build com golang:1.21-alpine
- [ ] Criar stage de runtime com alpine:latest
- [ ] Configurar exposição de porta
- [ ] Testar build

**Critérios de Aceite:**
- Docker build funciona
- Container executa corretamente

---

### Tarefa 6.3: Criar docker-compose.yml
**Status:** Pendente
**Prioridade:** Média
**Estimativa:** 0.25 dia

**Descrição:**
Criar docker-compose para facilitar deploy.

**Arquivo:** `docker-compose.yml`

**Subtarefas:**
- [ ] Configurar serviço api
- [ ] Configurar port mapping
- [ ] Configurar variáveis de ambiente
- [ ] Configurar restart policy

**Critérios de Aceite:**
- `docker-compose up` funciona
- Serviço acessível na porta 8080

---

## Fase 7: Testes Unitários

### Tarefa 7.1: Testes de Domínio
**Status:** Pendente
**Prioridade:** Média
**Estimativa:** 0.5 dia

**Descrição:**
Criar testes para entidades de domínio.

**Arquivos:** `*_test.go` na pasta domain

**Subtarefas:**
- [ ] Testar criação de Password
- [ ] Testar mascaramento de senha
- [ ] Testar cálculo de entropia
- [ ] Testar classificação de força

**Critérios de Aceite:**
- Todos os testes passam
- Cobertura ≥ 70% para domain

---

### Tarefa 7.2: Testes de Casos de Uso
**Status:** Pendente
**Prioridade:** Média
**Estimativa:** 0.5 dia

**Descrição:**
Criar testes para casos de uso.

**Arquivos:** `*_test.go` na pasta usecase

**Subtarefas:**
- [ ] Testar GeneratePassword
- [ ] Testar ValidatePassword
- [ ] Usar mocks para dependências
- [ ] Testar cenários de erro

**Critérios de Aceite:**
- Todos os testes passam
- Mocks funcionando

---

### Tarefa 7.3: Testes de Serviços
**Status:** Pendente
**Prioridade:** Média
**Estimativa:** 0.5 dia

**Descrição:**
Criar testes para serviços de infraestrutura.

**Arquivos:** `*_test.go` na pasta services

**Subtarefas:**
- [ ] Testar CryptoGenerator
- [ ] Testar RuleValidator
- [ ] Testar geração com diferentes configurações
- [ ] Testar validação com senhas inválidas

**Critérios de Aceite:**
- Todos os testes passam
- Cobertura ≥ 70% para services

---

### Tarefa 7.4: Testes de Handlers
**Status:** Pendente
**Prioridade:** Média
**Estimativa:** 0.5 dia

**Descrição:**
Criar testes para handlers HTTP.

**Arquivos:** `*_test.go` na pasta handlers

**Subtarefas:**
- [ ] Testar endpoint /generate
- [ ] Testar endpoint /validate
- [ ] Testar endpoint /health
- [ ] Testar tratamento de erros

**Critérios de Aceite:**
- Todos os testes passam
- Respostas HTTP corretas

---

## Fase 8: Documentação Swagger

### Tarefa 8.1: Adicionar Anotações Swagger
**Status:** Pendente
**Prioridade:** Baixa
**Estimativa:** 0.5 dia

**Descrição:**
Adicionar anotações Swagger nos handlers.

**Subtarefas:**
- [ ] Adicionar anotações no GeneratePassword
- [ ] Adicionar anotações no ValidatePassword
- [ ] Adicionar anotações no HealthCheck
- [ ] Definir schemas de request/response

**Critérios de Aceite:**
- Anotações adicionadas em todos os handlers

---

### Tarefa 8.2: Gerar Documentação
**Status:** Pendente
**Prioridade:** Baixa
**Estimativa:** 0.25 dia

**Descrição:**
Gerar documentação Swagger automaticamente.

**Subtarefas:**
- [ ] Executar `swag init -g cmd/api/main.go`
- [ ] Verificar geração de docs
- [ ] Integrar UI Swagger

**Critérios de Aceite:**
- Documentação gerada em `docs/`
- UI acessível em `/swagger/index.html`

---

## Fase 9: Coleção Postman

### Tarefa 9.1: Criar Collection Postman
**Status:** Pendente
**Prioridade:** Baixa
**Estimativa:** 0.25 dia

**Descrição:**
Criar collection Postman para testes da API.

**Arquivo:** `password_generator_collection.json`

**Subtarefas:**
- [ ] Criar request Generate Password
- [ ] Criar request Validate Password
- [ ] Configurar variáveis de ambiente
- [ ] Documentar responses

**Critérios de Aceite:**
- Collection importada no Postman
- Todos os requests funcionando

---

## Resumo de Tarefas

| Fase | Total Tarefas | Concluídas | Pendentes |
|------|---------------|------------|-----------|
| 1 | 4 | 4 | 0 |
| 2 | 2 | 2 | 0 |
| 3 | 2 | 2 | 0 |
| 4 | 2 | 2 | 0 |
| 5 | 3 | 3 | 0 |
| 6 | 3 | 3 | 0 |
| 7 | 4 | 4 | 0 |
| 8 | 2 | 2 | 0 |
| 9 | 1 | 1 | 0 |
| **Total** | **23** | **23** | **0** |