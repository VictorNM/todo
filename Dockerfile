FROM golang:1.12.9

WORKDIR /app

COPY . /app

RUN go build ./...

EXPOSE 80

CMD ./todo