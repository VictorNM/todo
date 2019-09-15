FROM golang:1.12.9

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 80

CMD ["./main"]