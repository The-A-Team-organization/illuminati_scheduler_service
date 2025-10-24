FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . . 
RUN go mod init service \
    && go get github.com/robfig/cron/v3 \
    && go get github.com/wneessen/go-mail \
    && go get github.com/XANi/loremipsum \
    && go get golang.org/x/crypto/bcrypt \
    && go mod tidy


RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
