FROM golang:1.25-alpine AS builder
WORKDIR /app
RUN go mod init service \
    &&go get github.com/robfig/cron/v3 \
    && go get github.com/wneessen/go-mail \
    && go get github.com/XANi/loremipsum \
    && go mod tidy

COPY . .

RUN go build -o main .
# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
