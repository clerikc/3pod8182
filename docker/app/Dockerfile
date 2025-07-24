# Сборка
FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o web-app .

# Запуск
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/web-app .
EXPOSE 8080
CMD ["./web-app"]