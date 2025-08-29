FROM golang:1.24-alpine AS builder

WORKDIR /app


RUN apk add --no-cache git build-base


ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=off


RUN go mod init delicious-and-kidney


RUN go get github.com/gin-gonic/gin@v1.10.1
RUN go get github.com/golang-jwt/jwt/v5@v5.3.0
RUN go get github.com/joho/godotenv@v1.5.1
RUN go get golang.org/x/crypto@v0.28.0
RUN go get gorm.io/driver/postgres@v1.6.0
RUN go get gorm.io/gorm@v1.30.1


COPY . .


ARG SERVICE_NAME


RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app ./cmd/${SERVICE_NAME}/


FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]