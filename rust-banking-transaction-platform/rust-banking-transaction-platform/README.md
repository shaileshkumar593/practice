# Rust Banking Transaction Platform

A production-oriented learning/reference project demonstrating a Rust backend using:

- Rust
- Tokio
- Axum
- Serde
- PostgreSQL
- Redis
- Kafka-ready transactional outbox
- Docker
- Kubernetes
- GCP deployment concepts
- OpenTelemetry/Prometheus integration points
- GitHub Actions CI

## Current implementation

Implemented:

- User creation and retrieval
- Account creation and retrieval
- Atomic account-to-account transfer
- PostgreSQL transaction
- Row-level locking with `FOR UPDATE`
- Deterministic account locking to reduce deadlocks
- Idempotency key persistence
- Transactional outbox
- Axum REST API
- Serde JSON
- SQLx PostgreSQL pool
- Redis client wiring
- Docker Compose
- Kubernetes Deployment/Service/HPA
- GitHub Actions CI
- Structured tracing

## Important production notes

This repository is an interview/learning reference, not a bank-ready production system.

Before production, add:

- authentication and authorization
- proper secret injection through GCP Secret Manager/Workload Identity
- Kafka outbox publisher and consumers
- Redis caching/rate limiter implementation
- OpenTelemetry exporter configuration
- Prometheus metrics endpoint/exporter
- stronger request validation
- integration/contract tests
- migration automation
- graceful shutdown
- retry/circuit-breaker policy where appropriate
- audit controls
- key management
- reconciliation
- double-entry ledger if required by the business model
- PCI/security controls where applicable

## Local run

Prerequisites:

- Rust stable
- Docker

Start PostgreSQL and Redis:

```bash
docker compose up -d postgres redis
```

Create the schema:

```bash
psql "postgres://banking:banking@localhost:5432/banking" -f migrations/0001_init.sql
```

Run:

```bash
cargo run
```

Health check:

```bash
curl http://localhost:3000/health
```

## Run everything with Docker

```bash
docker compose up --build
```

## Example API

Create a user:

```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@example.com"}'
```

Create an account:

```bash
curl -X POST http://localhost:3000/accounts \
  -H "Content-Type: application/json" \
  -d '{"user_id":"USER_UUID","currency":"INR"}'
```

Transfer:

```bash
curl -X POST http://localhost:3000/transfers \
  -H "Content-Type: application/json" \
  -d '{
    "from_account_id":"SOURCE_UUID",
    "to_account_id":"DESTINATION_UUID",
    "amount":500,
    "currency":"INR",
    "idempotency_key":"transfer-0001"
  }'
```

## Architecture

```text
Client
  |
GCP Load Balancer / Ingress
  |
GKE
  |
Axum + Tokio
  |
  +---- PostgreSQL
  |       |
  |       +---- business tables
  |       +---- idempotency keys
  |       +---- outbox events
  |
  +---- Redis
  |
  +---- Kafka (outbox publisher/consumers)
  |
OpenTelemetry / Prometheus
```

## Interview topics covered

- Rust ownership/borrowing
- async/await and Tokio
- Axum routing/extractors/state
- Serde
- PostgreSQL transactions and locking
- connection pooling
- idempotency
- transactional outbox
- Redis
- Kafka architecture
- Docker
- Kubernetes
- GCP/IAM/Secret Manager deployment concepts
- observability
- CI/CD
