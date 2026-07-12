# PRD - Gerador de Senhas com Validação

## 1. Visão Geral do Produto

### 1.1 Propósito
Sistema de geração e validação de senhas fortes com requisitos de complexidade, seguindo padrões de segurança para aplicações enterprise.

### 1.2 Objetivos de Negócio
- **Segurança**: Gerar senhas criptograficamente seguras usando `crypto/rand`
- **Conformidade**: Atender requisitos de senha (tamanho, caracteres especiais, entropia)
- **Usabilidade**: API REST intuitiva para integração com outros sistemas
- **Performance**: Processamento rápido sem reflexão ou anotações

### 1.3 Público-Alvo
- Desenvolvedores que precisam de serviço de geração de senhas
- Sistemas de autenticação e cadastro de usuários
- Aplicações que exigem validação de senha em tempo real

---

## 2. Requisitos Funcionais

### 2.1 Geração de Senha
| ID | Requisito | Prioridade | Detalhamento |
|----|-----------|------------|--------------|
| RF01 | Gerar senha aleatória | Alta | Usar `crypto/rand` para aleatoriedade segura |
| RF02 | Configurar comprimento | Alta | Parâmetro `length` (padrão: 12) |
| RF03 | Incluir maiúsculas | Média | Parâmetro `use_uppercase` (padrão: true) |
| RF04 | Incluir minúsculas | Média | Parâmetro `use_lowercase` (padrão: true) |
| RF05 | Incluir números | Média | Parâmetro `use_numbers` (padrão: true) |
| RF06 | Incluir símbolos | Média | Parâmetro `use_symbols` (padrão: true) |

### 2.2 Validação de Senha
| ID | Requisito | Prioridade | Detalhamento |
|----|-----------|------------|--------------|
| RF07 | Validar comprimento mínimo | Alta | Mínimo de 8 caracteres |
| RF08 | Validar comprimento máximo | Média | Máximo de 64 caracteres |
| RF09 | Validar caracteres especiais | Alta | Pelo menos 1 caractere especial |
| RF10 | Validar repetição consecutiva | Média | Máximo de 2 caracteres repetidos |
| RF11 | Calcular entropia | Alta | Medir força da senha em bits |
| RF12 | Classificar força | Alta | "weak" (<40), "medium" (40-59), "strong" (≥60) |

### 2.3 API REST
| ID | Requisito | Prioridade | Detalhamento |
|----|-----------|------------|--------------|
| RF13 | Endpoint `/generate` | Alta | POST /api/v1/password/generate |
| RF14 | Endpoint `/validate` | Alta | POST /api/v1/password/validate |
| RF15 | Retornar JSON | Alta | Respostas estruturadas em JSON |
| RF16 | Tratar erros | Alta | Mensagens claras de erro |
| RF17 | Health check | Média | GET /health |

### 2.4 Documentação e Testes
| ID | Requisito | Prioridade | Detalhamento |
|----|-----------|------------|--------------|
| RF18 | Documentação Swagger | Média | OpenAPI 3.0 |
| RF19 | Collection Postman | Média | JSON para testes dos endpoints |
| RF20 | Testes unitários | Alta | Cobertura mínima de 70% |

---

## 3. Requisitos Não Funcionais

### 3.1 Arquitetura
| ID | Requisito | Detalhamento |
|----|-----------|--------------|
| RNF01 | Clean Architecture | Separar em camadas: Entities, Use Cases, Interfaces, Frameworks |
| RNF02 | Modularidade | Pacotes bem definidos e com responsabilidades únicas |
| RNF03 | Stateless | Sem estado entre requisições, sem necessidade de banco de dados |

### 3.2 Performance
| ID | Requisito | Detalhamento |
|----|-----------|--------------|
| RNF04 | Tempo de resposta | < 50ms para geração e validação |
| RNF05 | Sem reflexão | Evitar `reflect` para manter performance |
| RNF06 | Concorrência | Suportar múltiplas requisições simultâneas |

### 3.3 Segurança
| ID | Requisito | Detalhamento |
|----|-----------|--------------|
| RNF07 | Aleatoriedade segura | Usar `crypto/rand`, nunca `math/rand` |
| RNF08 | Nunca logar senhas | Não registrar senhas em logs |
| RNF09 | CORS configurável | Permitir origens específicas |

### 3.4 DevOps
| ID | Requisito | Detalhamento |
|----|-----------|--------------|
| RNF10 | Dockerização | Dockerfile multi-stage |
| RNF11 | Configuração via env | Variáveis de ambiente para configuração |
| RNF12 | Portabilidade | Executar em qualquer sistema operacional |

---

## 4. Estrutura do Projeto

### 4.1 Clean Architecture - Organização das Camadas
password-generator/
├── cmd/
│ └── api/
│ └── main.go # Entry point
├── internal/
│ ├── domain/ # ENTITIES (Camada mais interna)
│ │ ├── password.go # Entidade Password
│ │ └── validation.go # Entidade ValidationResult
│ ├── usecase/ # USE CASES (Regras de negócio)
│ │ ├── generate_password.go # Caso de uso: Gerar senha
│ │ └── validate_password.go # Caso de uso: Validar senha
│ ├── interfaces/ # INTERFACES (Portas)
│ │ ├── password_generator.go # Interface do gerador
│ │ └── password_validator.go # Interface do validador
│ ├── infrastructure/ # FRAMEWORKS (Camada externa)
│ │ ├── handlers/ # Handlers HTTP (Gin)
│ │ │ └── password_handler.go
│ │ ├── services/ # Implementações concretas
│ │ │ ├── crypto_generator.go
│ │ │ └── rule_validator.go
│ │ ├── config/ # Configuração
│ │ │ └── config.go
│ │ └── router/ # Rotas
│ │ └── router.go
│ └── pkg/ # Pacotes compartilhados
│ └── errors/ # Erros customizados
└── go.mod

text

### 4.2 Fluxo da Clean Architecture
[Requisição HTTP]
↓
[Handlers - Camada Frameworks]
↓
[Use Cases - Camada Use Cases]
↓
[Domain Entities - Camada Entities]
↓
[Services - Camada Infrastructure]
↓
[Resposta HTTP]

text

### 4.3 Dependências
| Módulo | Versão | Propósito |
|--------|--------|-----------|
| `github.com/gin-gonic/gin` | v1.9.0 | Framework HTTP |
| `github.com/spf13/viper` | v1.16.0 | Configuração |
| `github.com/stretchr/testify` | v1.8.4 | Testes |
| `github.com/swaggo/swag` | v1.16.1 | Documentação Swagger |
| `github.com/swaggo/gin-swagger` | v1.6.0 | UI Swagger |
| `golang.org/x/crypto` | v0.12.0 | Funções criptográficas |

---

## 5. Endpoints da API

### 5.1 Gerar Senha
POST /api/v1/password/generate
Content-Type: application/json

{
"length": 16,
"use_uppercase": true,
"use_lowercase": true,
"use_numbers": true,
"use_symbols": true
}

Response (200 OK):
{
"password": "xH7@kL9#mP2&",
"validation": {
"is_valid": true,
"errors": [],
"strength": "strong",
"entropy": 72.54
},
"generated": "2026-07-11T14:30:00Z"
}

text

### 5.2 Validar Senha
POST /api/v1/password/validate
Content-Type: application/json

{
"password": "xH7@kL9#mP2&"
}

Response (200 OK):
{
"is_valid": true,
"errors": [],
"strength": "strong",
"entropy": 72.54
}

text

### 5.3 Health Check
GET /health

Response (200 OK):
{
"status": "healthy",
"timestamp": "2026-07-11T14:30:00Z"
}

text

---

## 6. Dockerização

### 6.1 Dockerfile Multi-stage
```dockerfile
# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /password-api ./cmd/api

# Stage 2: Runtime
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /password-api .
EXPOSE 8080
CMD ["./password-api"]
6.2 docker-compose.yml
yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - MIN_LENGTH=8
      - MAX_LENGTH=64
      - MIN_SPECIAL_CHARS=1
      - MIN_ENTROPY=40
    restart: unless-stopped
7. Configuração via Variáveis de Ambiente
Variável	Padrão	Descrição
PORT	8080	Porta da API
MIN_LENGTH	8	Comprimento mínimo da senha
MAX_LENGTH	64	Comprimento máximo da senha
MIN_SPECIAL_CHARS	1	Mínimo de caracteres especiais
MIN_ENTROPY	40	Entropia mínima (bits)
MAX_CONSECUTIVE	2	Máximo de repetições consecutivas
8. Coleção de Testes (Postman)
Arquivo: password_generator_collection.json

json
{
  "info": {
    "name": "Password Generator API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Generate Password",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\"length\":16,\"use_uppercase\":true,\"use_lowercase\":true,\"use_numbers\":true,\"use_symbols\":true}"
        },
        "url": {
          "raw": "http://localhost:8080/api/v1/password/generate",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["api", "v1", "password", "generate"]
        }
      }
    },
    {
      "name": "Validate Password",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\"password\":\"xH7@kL9#mP2&\"}"
        },
        "url": {
          "raw": "http://localhost:8080/api/v1/password/validate",
          "protocol": "http",
          "host": ["localhost"],
          "port": "8080",
          "path": ["api", "v1", "password", "validate"]
        }
      }
    }
  ]
}
9. Plano de Implementação
Fase	Tarefas	Estimativa	Prioridade
1	Setup do projeto e estrutura de pastas	1 dia	Alta
2	Implementação da camada Domain	1 dia	Alta
3	Implementação dos Use Cases	2 dias	Alta
4	Implementação dos Services	2 dias	Alta
5	Implementação dos Handlers e Router	1 dia	Alta
6	Configuração e Dockerização	1 dia	Média
7	Testes unitários	2 dias	Média
8	Documentação Swagger	1 dia	Baixa
9	Coleção Postman	0.5 dia	Baixa
Total estimado: 11.5 dias

10. Critérios de Aceite
10.1 Funcionais
Gerar senha com crypto/rand em < 10ms

Validar todos os requisitos de complexidade

Retornar classificação de força correta

Endpoints respondem em < 50ms

Container Docker sobe sem erros

10.2 Técnicos
Código segue Clean Architecture

Cobertura de testes ≥ 70%

Sem uso de reflect ou unsafe

Logs estruturados e sem vazamento de senhas

Documentação Swagger acessível em /swagger/index.html

10.3 Segurança
Uso exclusivo de crypto/rand

Senhas nunca aparecem em logs

CORS configurável

11. Métricas de Sucesso
Métrica	Alvo	Como Medir
Tempo de resposta	< 50ms	Testes de carga
Uptime	99.9%	Monitoramento
Entropia média	≥ 50 bits	Dashboard
Satisfação do dev	≥ 4.5/5	Pesquisa
12. Riscos e Mitigações
Risco	Probabilidade	Impacto	Mitigação
crypto/rand falhar	Baixa	Alto	Retry mechanism
Performance em alto volume	Média	Médio	Implementar pool de workers
Vazamento de senha em logs	Baixa	Crítico	Revisão de código obrigatória
13. Diferenciais Técnicos
13.1 Go vs Java/Spring Security
Aspecto	Go	Java (Spring)
Aleatoriedade	crypto/rand nativo	SecureRandom
Validação	Funções puras com regexp	Anotações @Pattern, @Size
Performance	Sem reflexão (rápido)	Usa reflexão (mais lento)
Complexidade	Estruturas explícitas	Anotações e interfaces
Memória	Baixo overhead	Alto consumo
13.2 Por que Clean Architecture em Go?
Separação clara de responsabilidades

Facilita testes unitários com mocks

Permite trocar frameworks (ex: Gin → Echo) sem impacto

Código mais maintainable e escalável

14. Comandos Úteis (Desenvolvimento)
bash
# Executar aplicação
go run cmd/api/main.go

# Build da aplicação
go build -o password-api ./cmd/api

# Executar testes
go test ./... -v -cover

# Rodar com Docker
docker build -t password-api .
docker run -p 8080:8080 password-api

# Rodar com docker-compose
docker-compose up -d

# Gerar documentação Swagger
swag init -g cmd/api/main.go

# Ver dependências
go mod tidy
go mod verify
15. Referências
Go Crypto Package

Clean Architecture (Robert Martin)

Gin Framework

OWASP Password Guidelines

Entropy in Passwords