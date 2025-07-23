FROM golang:1.21-alpine
WORKDIR /app
COPY main.go .
RUN go build -o web-app .
EXPOSE 8080
CMD ["./web-app"]