FROM golang:1.21-alpine

WORKDIR /app

# Копируем весь проект
COPY . .

# Инициализируем go модуль
RUN go mod init delicious-and-kidney

# Подтягиваем зависимости
RUN go mod tidy

# Собираем приложение
RUN go build -o auth-service ./cmd/auth-service/

EXPOSE 8080
CMD ["./auth-service"]