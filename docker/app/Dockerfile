FROM golang:1.24.4-alpine

WORKDIR /app

# Копируем сначала go.mod и go.sum (если есть)
COPY go.mod .
RUN go mod download

# Копируем остальные файлы
COPY main.go .

# Собираем приложение
RUN go build -o web-app .

EXPOSE 8080
CMD ["./web-app"]