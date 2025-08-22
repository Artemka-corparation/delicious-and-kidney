FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG SERVICE_NAME


RUN go build -o app ./cmd/${SERVICE_NAME}/

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]