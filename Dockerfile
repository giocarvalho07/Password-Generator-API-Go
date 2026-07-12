# Estágio de Build
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Estágio de Execução (Seguro para OpenShift)
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
# Garante que o binário tem permissão de execução
RUN chmod +x ./main

# Portas abaixo de 1024 exigem privilégios de root, use 8080
EXPOSE 8080
CMD ["./main"]