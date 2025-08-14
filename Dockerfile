FROM golang:1.23-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/server/main.go

EXPOSE 8080

CMD ["./main"]