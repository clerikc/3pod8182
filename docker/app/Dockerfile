# Сборка
FROM golang:1.24-alpine AS builder
WORKDIR /app
# Копируем go.mod из корня проекта
COPY ./docker/app/go.mod .
RUN go mod download
# Копируем main.go из docker/app
COPY ./docker/app/main.go .
RUN CGO_ENABLED=0 go build -o web-app .

# Запуск
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/web-app .
EXPOSE 8080
CMD ["./web-app"]