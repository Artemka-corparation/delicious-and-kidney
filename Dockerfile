FROM golang:1.24.3-alpine

WORKDIR /app

COPY go.mod go.sum

COPY . .

RUN go mod tidy && go build -o main ./cmd/auth-service

EXPOSE 8080

CMD ["./main"]