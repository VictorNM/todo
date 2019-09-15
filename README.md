# GOLANG TODO APP

## Build and Run

```bash
go build ./...

./todo

# Test
curl http://localhost:8080 # should not error
```

## Docker

```bash
docker build --tag=todo:lastest .
docker run -d -p 8080:8080 todo

# Test
curl http://localhost:8080 # should not error
```

- Referrences: 
    + [https://docs.docker.com/get-started/part2/](https://docs.docker.com/get-started/part2/)
    + [https://hub.docker.com/_/golang](https://hub.docker.com/_/golang)
    + [https://www.callicoder.com/docker-golang-image-container-example/](https://www.callicoder.com/docker-golang-image-container-example/)