# Golang Layered Architecture CRUD API

## Architecture

- Handler Layer
- Service Layer
- Repository Layer
- Model Layer

Each layer is fully decoupled using interfaces.

---

## Features

- Concurrent-safe map using sync.RWMutex
- GET
- POST
- PUT
- DELETE
- Dependency Injection
- Interface-based architecture

---

## Run Project

```bash
go run ./cmd/server
```

---

## Endpoints

### Create Employee

```bash
curl -X POST http://localhost:8080/employees \
-H "Content-Type: application/json" \
-d '{"id":"1","name":"Shailesh"}'
```

### Get All

```bash
curl http://localhost:8080/employees
```

### Get By ID

```bash
curl http://localhost:8080/employees/1
```

### Update

```bash
curl -X PUT http://localhost:8080/employees/1 \
-H "Content-Type: application/json" \
-d '{"name":"Kumar"}'
```

### Delete

```bash
curl -X DELETE http://localhost:8080/employees/1
```
