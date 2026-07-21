# Fluxo do Projeto - Password Validate

## Sequência Completa (Request → Response)

```
┌─────────────────────────────────────────────────────────────────┐
│                     CLIENTE (HTTP Request)                      │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  1. ROUTER                                                      │
│     internal/infrastructure/router/router.go                    │
│     • Recebe requisição HTTP                                    │
│     • Verifica método (GET/POST) e rota                         │
│     • Encaminha para o handler correto                          │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  2. HANDLER                                                     │
│     internal/infrastructure/handlers/password_handler.go        │
│     • Parse do body JSON (bind)                                 │
│     • Validação básica do request                               │
│     • Chama o UseCase                                           │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  3. USE CASE                                                    │
│     internal/usecase/generate_password.go                       │
│     internal/usecase/validate_password.go                       │
│     • Aplica regras de negócio                                  │
│     • Delega para a Interface (contrato)                        │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  4. INTERFACE                                                   │
│     internal/interfaces/password_generator.go                   │
│     internal/interfaces/password_validator.go                   │
│     • Contrato que define o comportamento                       │
│     • Separa dependências (Inversão de Dependência)             │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  5. SERVICE (Implementação Concreta)                            │
│     internal/infrastructure/services/crypto_generator.go        │
│     internal/infrastructure/services/rule_validator.go          │
│     • CryptoGenerator: usa crypto/rand para gerar senha        │
│     • RuleValidator: aplica regras de validação                 │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  6. DOMAIN                                                      │
│     internal/domain/password.go                                 │
│     internal/domain/validation.go                               │
│     • Entidade Password: mascaração, detecção de caracteres     │
│     • Entidade ValidationResult: entropia, classificação        │
│     • Cálculo: E = L × log2(N)                                  │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│  7. RESPOSTA (Response)                                         │
│     Handler retorna JSON ao cliente                             │
└─────────────────────────────────────────────────────────────────┘
```

---

## Fluxo Invertido (Response → Request)

```
┌─────────────────────────────────────────────────────────────────┐
│  7. DOMAIN                                                      │
│     Resultado calculado (senha, validação, entropia)            │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▲
┌─────────────────────────────────────────────────────────────────┐
│  6. SERVICE                                                     │
│     Service retorna resultado processado                        │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▲
┌─────────────────────────────────────────────────────────────────┐
│  5. INTERFACE                                                   │
│     Interface recebe resultado via contrato                     │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▲
┌─────────────────────────────────────────────────────────────────┐
│  4. USE CASE                                                    │
│     UseCase formata resposta de negócio                         │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▲
┌─────────────────────────────────────────────────────────────────┐
│  3. HANDLER                                                     │
│     Handler monta response JSON                                │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▲
┌─────────────────────────────────────────────────────────────────┐
│  2. ROUTER                                                      │
│     Router envia resposta HTTP                                  │
└─────────────────────────────────────────────────────────────────┘
                                │
                                ▲
┌─────────────────────────────────────────────────────────────────┐
│  1. CLIENTE (HTTP Response)                                     │
└─────────────────────────────────────────────────────────────────┘
```

---

## Rotas Disponíveis

| Método | Rota | Descrição |
|--------|------|-----------|
| `GET` | `/health` | Health check |
| `POST` | `/api/v1/password/generate` | Gerar senha |
| `POST` | `/api/v1/password/validate` | Validar senha |
| `GET` | `/swagger/*any` | Documentação Swagger |

---

## Fluxo por Endpoint

### POST /api/v1/password/generate

```
Request → Router → Handler → UseCase(Generate) 
  → Interface(Generator) → Service(CryptoGenerator) 
  → Domain(Password) → Response(senha + validação)
```

### POST /api/v1/password/validate

```
Request → Router → Handler → UseCase(Validate) 
  → Interface(Validator) → Service(RuleValidator) 
  → Domain(ValidationResult) → Response(resultado)
```

### GET /health

```
Request → Router → Handler(HealthCheck) → Response(status)
```

---

## Camadas da Clean Architecture

```
┌─────────────────────────────────────────┐
│          INFRASTRUCTURE                 │
│  (Handlers, Services, Config, Router)   │
├─────────────────────────────────────────┤
│          INTERFACES                     │
│  (Contratos: Generator, Validator)      │
├─────────────────────────────────────────┤
│          USE CASES                      │
│  (Regras de Negócio)                    │
├─────────────────────────────────────────┤
│          DOMAIN                         │
│  (Entidades: Password, Validation)      │
└─────────────────────────────────────────┘
```

**Regra**: Cada camada só depende da camada abaixo (dependência para baixo).
