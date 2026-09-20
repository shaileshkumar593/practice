# Axum Layered User API

Simple Rust/Axum CRUD API using layered architecture.

Request flow:

HTTP -> Routes -> Handlers -> Service -> Repository -> In-memory storage

## Structure

src/
├── main.rs
├── config.rs
├── routes/
│   └── user.rs
├── handlers/
│   └── user.rs
├── services/
│   └── user_service.rs
├── repositories/
│   └── user_repository.rs
├── models/
│   └── user.rs
├── errors/
│   └── app_error.rs
└── state.rs

The `mod.rs` files are included to expose each module.

## Run

```bash
cargo run
```

Server: http://localhost:3000

## Endpoints

POST /users
GET /users/{id}
PUT /users/{id}
DELETE /users/{id}

Create:

```bash
curl -X POST http://localhost:3000/users ^
  -H "Content-Type: application/json" ^
  -d "{"name":"Alice","email":"alice@example.com"}"
```

Get:

```bash
curl http://localhost:3000/users/<UUID>
```

Update:

```bash
curl -X PUT http://localhost:3000/users/<UUID> ^
  -H "Content-Type: application/json" ^
  -d "{"name":"Alice Updated","email":"alice.updated@example.com"}"
```

Delete:

```bash
curl -X DELETE http://localhost:3000/users/<UUID>
```

## Notes

The repository uses an in-memory `HashMap` for simplicity. In a production service, replace it with PostgreSQL/SQLx while preserving the route -> handler -> service -> repository separation.
