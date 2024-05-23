# Usar una imagen base de Go más ligera
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o backend-app

# Usar una imagen base mínima para ejecutar la aplicación
FROM alpine:latest

WORKDIR /app

# Copiar el ejecutable desde la etapa de construcción
COPY --from=builder /app/backend-app .

EXPOSE 8080

CMD ["./backend-app", ":8080"]