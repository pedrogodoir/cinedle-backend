# Etapa 1: Build
FROM golang:1.25.1 AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos de dependência primeiro (para cache eficiente)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário (sem dependência de C)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api/main.go

# Etapa 2: Imagem final (produção)
FROM debian:bullseye-slim

# Instala certificados SSL (para conexões HTTPS e com o banco Neon)
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Define o diretório de trabalho
WORKDIR /root/

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/server .

# Expõe a porta da API
EXPOSE 8080

# Comando para rodar o servidor
CMD ["./server"]
