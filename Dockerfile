FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o main ./cmd/main.go

EXPOSE 80

CMD ["./main"]