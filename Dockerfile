# =========================
# 1️⃣ Stage de build
# =========================
FROM golang:1.21-alpine AS builder

# Instala dependências do sistema
RUN apk add --no-cache git

# Define diretório de trabalho
WORKDIR /app

# Copia arquivos de dependências para cache
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código da aplicação
COPY . .

# Instala o swag CLI para gerar documentação
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Gera a documentação Swagger
RUN swag init -g app/servidor/main.go

# Compila a aplicação Go
RUN go build -o server ./app/servidor/main.go

# =========================
# 2️⃣ Stage final mínimo
# =========================
FROM alpine:latest

# Define diretório de trabalho
WORKDIR /app

# Copia binário e documentação gerada do builder
COPY --from=builder /app/server .
COPY --from=builder /app/docs ./docs

# Expõe a porta do serviço
EXPOSE 8081

# Comando para rodar a aplicação
CMD ["./server"]
