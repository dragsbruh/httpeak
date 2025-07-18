FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -ldflags="-s -w" -o httpeak main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/httpeak .

EXPOSE 80

CMD ["./httpeak"]