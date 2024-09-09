# # Etapa 1: Construção da aplicação
# FROM golang:1.23 AS builder

# # Definir o diretório de trabalho
# WORKDIR /app

# # Copiar o go.mod e go.sum e baixar as dependências
# COPY go.mod go.sum ./
# RUN go mod download

# # Copiar o código-fonte
# COPY . .

# # Construir a aplicação
# RUN go build -o app

# # Etapa 2: Execução da aplicação
# FROM debian:bullseye-slim

# # Instalar dependências necessárias
# RUN apt-get update && apt-get install -y \
#     ca-certificates \
#     && rm -rf /var/lib/apt/lists/*

# # Definir o diretório de trabalho
# WORKDIR /root/

# # Copiar o binário da etapa de construção
# COPY --from=builder /app/app .

# # Expor a porta que o servidor vai usar
# EXPOSE 8080

# # Comando para rodar a aplicação
# CMD ["./app"]

# Etapa 1: Construção da aplicação
# Etapa 1: Usar a imagem oficial do Golang para rodar o código
FROM golang:1.23

# Definir o diretório de trabalho
WORKDIR /app

# Copiar o go.mod e go.sum e baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código-fonte para o container
COPY . .

# Expor a porta que o servidor vai usar
EXPOSE 8080

# Comando para rodar a aplicação (usando 'go run')
CMD ["go", "run", "main.go"]
