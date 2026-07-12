# Especificação Técnica - Gerador de Senhas com Validação

## 1. Visão Geral do Sistema

### 1.1 Arquitetura
Sistema desenvolvido em Go seguindo Clean Architecture com as seguintes camadas:
- **Domain**: Entidades de negócio (Password, ValidationResult)
- **Use Cases**: Regras de geração e validação de senhas
- **Interfaces**: Portas para comunicação entre camadas
- **Infrastructure**: Implementações concretas (handlers HTTP, serviços de crypto)

### 1.2 Stack Tecnológica
| Componente | Tecnologia | Versão |
|------------|------------|--------|
| Linguagem | Go | 1.21 |
| Framework HTTP | Gin | v1.9.0 |
| Configuração | Viper | v1.16.0 |
| Testes | Testify | v1.8.4 |
| Documentação | Swag | v1.16.1 |
| Criptografia | crypto/rand | nativo |

## 2. Especificação dos Componentes

### 2.1 Entidades de Domínio

#### Password
```go
type Password struct {
    Value     string
    Length    int
    Generated time.Time
}
```

#### ValidationResult
```go
type ValidationResult struct {
    IsValid  bool
    Errors   []string
    Strength string
    Entropy  float64
}
```

### 2.2 Casos de Uso

#### GeneratePassword
- **Entrada**: Configuração de geração (tamanho, tipos de caracteres)
- **Saída**: Senha gerada + Validação
- **Regra**: Usar crypto/rand para geração segura

#### ValidatePassword
- **Entrada**: Senha para validação
- **Saída**: ValidationResult com força e entropia
- **Regras**:
  - Comprimento: 8-64 caracteres
  - Pelo menos 1 caractere especial
  - Máximo 2 caracteres repetidos consecutivos
  - Cálculo de entropia em bits

### 2.3 Interfaces

#### PasswordGenerator
```go
type PasswordGenerator interface {
    Generate(config GenerateConfig) (string, error)
}
```

#### PasswordValidator
```go
type PasswordValidator interface {
    Validate(password string) (*ValidationResult, error)
}
```

### 2.4 Serviços de Infraestrutura

#### CryptoGenerator
- Implementação de PasswordGenerator usando crypto/rand
- Suporte a caracteres maiúsculos, minúsculos, números e símbolos

#### RuleValidator
- Implementação de PasswordValidator
- Validação de regras de complexidade
- Cálculo de entropia usando Shannon entropy

## 3. Especificação da API

### 3.1 Endpoints

#### POST /api/v1/password/generate
**Request:**
```json
{
  "length": 16,
  "use_uppercase": true,
  "use_lowercase": true,
  "use_numbers": true,
  "use_symbols": true
}
```

**Response 200:**
```json
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
```

#### POST /api/v1/password/validate
**Request:**
```json
{
  "password": "xH7@kL9#mP2&"
}
```

**Response 200:**
```json
{
  "is_valid": true,
  "errors": [],
  "strength": "strong",
  "entropy": 72.54
}
```

#### GET /health
**Response 200:**
```json
{
  "status": "healthy",
  "timestamp": "2026-07-11T14:30:00Z"
}
```

### 3.2 Códigos de Resposta
| Código | Descrição |
|--------|-----------|
| 200 | Sucesso |
| 400 | Requisição inválida |
| 422 | Senha não atende requisitos |
| 500 | Erro interno do servidor |

## 4. Especificação de Segurança

### 4.1 Requisitos de Segurança
- Uso exclusivo de `crypto/rand` (nunca `math/rand`)
- Senhas nunca devem aparecer em logs
- CORS configurável para origens específicas
- Validação de entrada em todos os endpoints

### 4.2 Cálculo de Entropia
- Fórmula: `E = L * log2(N)`
- Onde:
  - L = comprimento da senha
  - N = tamanho do alfabeto utilizado
- Classificação:
  - < 40 bits: weak
  - 40-59 bits: medium
  - ≥ 60 bits: strong

## 5. Especificação de Configuração

### 5.1 Variáveis de Ambiente
| Variável | Padrão | Descrição |
|----------|--------|-----------|
| PORT | 8080 | Porta da API |
| MIN_LENGTH | 8 | Comprimento mínimo |
| MAX_LENGTH | 64 | Comprimento máximo |
| MIN_SPECIAL_CHARS | 1 | Mínimo de caracteres especiais |
| MIN_ENTROPY | 40 | Entropia mínima (bits) |
| MAX_CONSECUTIVE | 2 | Máximo de repetições consecutivas |

## 6. Especificação de Testes

### 6.1 Critérios de Cobertura
- Cobertura mínima de 70%
- Testes unitários para cada componente
- Testes de integração para endpoints

### 6.2 Cenários de Teste
1. Geração de senha com configurações padrão
2. Geração de senha com configurações personalizadas
3. Validação de senha válida
4. Validação de senha inválida (diversos cenários)
5. Cálculo correto de entropia
6. Tratamento de erros
7. Health check

## 7. Especificação de Deploy

### 7.1 Docker
- Multi-stage build (builder + runtime)
- Imagem base: alpine:latest
- Exposição da porta 8080

### 7.2 docker-compose
- Serviço único: api
- Port mapping: 8080:8080
- Variáveis de ambiente configuráveis
- Restart policy: unless-stopped