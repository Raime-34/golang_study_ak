# Используем официальный образ Go
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY course3/2.server/2.server_http/task3.2.2.2 .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# Используем минимальный образ для выполнения
FROM alpine:latest

# Копируем бинарник из предыдущего этапа
COPY --from=builder /app/server /app/server

# Указываем рабочую директорию
WORKDIR /app

# Экспортируем порт
EXPOSE 8080

RUN chmod +x /app/server

# Указываем команду для запуска
CMD ["./server"]
