# GOLANG TODO APP

## Build and Run

```bash
go build -o main ./cmd/main.go

./main

# Test
curl http://localhost:8080 # should not error
```

## Docker

```bash
docker build --tag=todo:latest .
docker run --name=todo -d -p 8080:8080 todo

# Log
docker logs -f todo

# Test
curl http://localhost:8080 # should not error

# Stop
docker container stop todo
```

- References: 
    + [https://docs.docker.com/get-started/part2/](https://docs.docker.com/get-started/part2/)
    + [https://hub.docker.com/_/golang](https://hub.docker.com/_/golang)
    + [https://www.callicoder.com/docker-golang-image-container-example/](https://www.callicoder.com/docker-golang-image-container-example/)