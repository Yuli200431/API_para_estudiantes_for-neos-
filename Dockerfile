# ---- Etapa 1: compilar ----
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copiar dependencias primero (mejor uso de caché)
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar el binario
RUN go build -o api ./cmd/api/main.go

# ---- Etapa 2: imagen final ----
FROM alpine:3.19

WORKDIR /app

# Copiar solo el binario compilado
COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]