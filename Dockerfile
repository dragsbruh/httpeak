FROM golang:1.22-alpine

WORKDIR /app

COPY main.go .

RUN go build -o httpeak main.go

EXPOSE 80

CMD ["./httpeak"]