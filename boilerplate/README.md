# Go-Kit Production-Style API Boilerplate

A small interview-friendly Go-Kit service demonstrating:

- Project structure
- Configuration via environment variables
- PostgreSQL connection
- Model/entity
- Repository layer
- Service layer
- Go-Kit endpoint layer
- HTTP transport/router
- Error handling middleware
- Structured logging with Zap
- Unit tests
- SQL migration
- Health endpoint
- Docker Compose for PostgreSQL
- Makefile

## Architecture

```text
Client
  |
  v
HTTP Router / Go-Kit Transport
  |
  v
Endpoint
  |
  v
Service
  |
  v
Repository
  |
  v
PostgreSQL
```

## Request flow

```text
GET /v1/users/1
       |
       v
HTTP Handler
       |
       v
Decode Request
       |
       v
Go-Kit Endpoint
       |
       v
UserService.GetUser
       |
       v
PostgresUserRepository.GetByID
       |
       v
PostgreSQL
       |
       v
Encode Response
       |
       v
HTTP 200 JSON
```

## Project structure

```text
.
├── cmd/api/main.go
├── internal/
│   ├── config/config.go
│   ├── database/postgres.go
│   ├── endpoint/user_endpoint.go
│   ├── middleware/logging.go
│   ├── model/user.go
│   ├── repository/user_repository.go
│   ├── service/user_service.go
│   └── transport/http/user_handler.go
├── migrations/
│   ├── 000001_create_users.up.sql
│   └── 000001_create_users.down.sql
├── tests/
│   └── user_service_test.go
├── docker-compose.yml
├── .env.example
├── Makefile
└── go.mod
```

## Prerequisites

- Go 1.23+
- PostgreSQL 14+
- Optional: Docker + Docker Compose
- Optional: golang-migrate CLI

## Run with Docker PostgreSQL

```bash
cp .env.example .env
docker compose up -d 
docker compose up -d --build 
```

Run migration:

```bash
migrate -path migrations \
  -database "postgres://postgres:postgres@localhost:5432/appdb?sslmode=disable" up

  go run ./cmd/migrate -command=up   
```

Then:

```bash
go mod tidy
go run ./cmd/api
```

API:

```text
GET http://localhost:8080/health
GET http://localhost:8080/v1/users/1
POST http://localhost:8080/v1/users
```

Create a user:

```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'
```

Get a user:

```bash
curl http://localhost:8080/v1/users/1
```

## Tests

```bash
go test ./...
go test ./... -cover
```

The service tests use a fake repository, so unit tests do not require PostgreSQL.

## Why Go-Kit?

Go-Kit separates transport, endpoint, and business logic. This makes services easier to test and lets the same business endpoint be exposed through HTTP, gRPC, messaging, or other transports.

## Production extensions

For a production service, add:

- Request ID / correlation ID
- OpenTelemetry tracing
- Prometheus metrics
- JWT/OAuth2 authentication
- Rate limiting
- Graceful shutdown
- Connection pool tuning
- Retry/circuit breaker where appropriate
- gRPC transport
- Kafka/event publishing
- Outbox pattern for reliable events
- Kubernetes manifests / Helm
- CI/CD



Understanding of SSO & MFA
✅ Experience with one or more of:
🔹 OAuth 2.0 / OIDC
🔹 SAML
🔹 LDAP / Active Directory
🔹 Kerberos
🔹 PKI / X.509
🔹 FIDO2 / WebAuthn
Bazel & Monorepo
🔹 OSS Package Upgrades