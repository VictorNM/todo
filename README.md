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
docker run -d -p 8080:8080 todo

# Log
docker logs -f {ID}

# Test
curl http://localhost:8080 # should not error

# Stop
docker container stop {ID}
```

- References: 
    + [https://docs.docker.com/get-started/part2/](https://docs.docker.com/get-started/part2/)
    + [https://hub.docker.com/_/golang](https://hub.docker.com/_/golang)
    + [https://www.callicoder.com/docker-golang-image-container-example/](https://www.callicoder.com/docker-golang-image-container-example/)
    
## APIs

### Get all Todos:

- Method: `GET`
    
- URL: `/todos`

- URL params: `none`

- Body: `none`

- Example:
    
    `GET/todos`
    
- Success response:
    - Code: `200`
    - Content:
    ```json
    {
        "data": [
            {
                "id": 1,
                "title": "Hello",
                "text": "World",
                "complete": false
            }
        ],
        "error": null
    }
    ```
  
- Error response:
    - Code: `400`
    - Content:
    ```json
    {
        "data": null,
        "error": {}
    }
    ```
    
### Get one Todo by ID:

- Method: `GET`
    
- URL: `/todos/:id`

- URL params:

    `id` : `integer` `required`

- Body: `none`

- Example:
    
    `GET` `/todos/1`
    
- Success response:
    - Code: `200`
    - Content:
    ```json
    {
        "data": {
            "id": 1,
            "title": "Hello",
            "text": "World",
            "complete": false
        },
        "error": null
    }
    ```
  
- Error response:
    - Code: `404`
    - Content:
    ```json
    {
        "data": null,
        "error": {}
    }
    ```
  
### Create new Todo:

- Method: `POST`
    
- URL: `/todos`

- URL params: `none`

- Body:

    `title` : `string` `required`
    
    `text` : `string` `required`

- Example:
    
    `POST` `/todos`
    ```json
    {
        "title": "Hello",
        "text": "World"
    }
    ```
    
- Success response:
    - Code: `201`
    - Content:
    ```json
        {
            "data": {
                "id": 1,
                "title": "Hello",
                "text": "World",
                "complete": false
            },
            "error": null
        }
    ```
  
- Error response:
    - Code: `400`
    - Content:
    ```json
    {
        "data": null,
        "error": {}
    }
    ```

### Update or create a Todo:

- Method: `PUT`
    
- URL: `/todos/:id`

- URL params:

    `id` : `integer` `required`

- Body:

    `title` : `string`
    
    `text` : `string`
    
    `complete` : `boolean` 

- Example:
    
    `PUT` `/todos/1`
    ```json
    {
        "title": "Hello",
        "text": "World",
        "complete": true
    }
    ```

- Success response:
    - Code: `200`
    - Content:
    ```json
    {
        "data": [
            {
                "id": 1,
                "title": "Hello",
                "text": "World",
                "complete": true
            }
        ],
        "error": null
    }
    ```
  
- Error response:
    - Code: `400`
    - Content:
    ```json
    {
        "data": null,
        "error": {}
    }
    ```
  
### Delete a Todo:

- Method: `DELETE`
    
- URL: `/todos/:id`

- URL params:

    `id` : `integer` `required`

- Body: `none`

- Example:
    
    `DELETE` `/todos/1`
    
- Success response:
    - Code: `204`
    - Content: `none`
  
- Error response:
    - Code: `400`
    - Content:
    ```json
    {
        "data": null,
        "error": {}
    }
    ```