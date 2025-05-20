# Этап сборки
FROM golang:latest AS builder
WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 go build -o /app/server ./src/cmd/

# Финальный этап
FROM alpine:latest
WORKDIR /app

# Устанавливаем CA сертификаты и часовые пояса
RUN apk --no-cache add ca-certificates tzdata && \
    update-ca-certificates

# Копируем бинарный файл из этапа сборки
COPY --from=builder /app/server /app/server

# Копируем необходимые директории и файлы
COPY --from=builder /app/frontend/dist /app/frontend/dist
COPY --from=builder /app/public /app/public
COPY --from=builder /app/docs /app/docs
#COPY --from=builder /app/.env /app/.env

# Устанавливаем рабочую директорию и экспонируем порт
WORKDIR /app
