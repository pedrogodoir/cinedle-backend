# Etapa de build
FROM golang:1.25 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/api

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 5000
CMD ["./main"]
