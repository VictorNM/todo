FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 80

CMD ["./main"]