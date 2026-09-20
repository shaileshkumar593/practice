# Rust Backend Interview Project

Production-style learning project covering:

- Rust ownership, borrowing, lifetimes, traits, generics, enums and error handling
- Tokio async runtime
- Axum REST API
- Serde JSON
- PostgreSQL + transactions + outbox
- Redis caching + TTL
- Kafka producer
- OpenTelemetry-ready tracing architecture
- Prometheus metrics endpoint
- Docker multi-stage build
- Kubernetes deployment, probes and HPA
- GCP deployment concepts
- IAM and Secret Manager integration points
- GitHub Actions CI/CD

## Run locally

1. Start infrastructure:

```bash
docker compose up -d
```

2. Copy environment:

```bash
cp .env.example .env
```

3. Run:

```bash
cargo run
```

4. Test:

```bash
curl http://localhost:3000/health
curl http://localhost:3000/ready
curl http://localhost:3000/metrics
```

## Create user

```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'
```

## Get user

```bash
curl http://localhost:3000/api/v1/users/<UUID>
```

## Important production notes

This repository is an interview/learning reference. Before production use:

- Use SQLx migrations instead of runtime table creation.
- Add a real Kafka outbox publisher/relay.
- Add authentication middleware with JWT validation.
- Store secrets in GCP Secret Manager, not Kubernetes YAML.
- Use Workload Identity on GKE.
- Add OpenTelemetry exporters and a trace backend.
- Add Kafka consumer groups, retries and DLQ handling.
- Add idempotency keys for money transfers.
- Add authorization/RBAC.
- Add integration/contract/load tests.
- Pin dependency versions and scan dependencies/images.
