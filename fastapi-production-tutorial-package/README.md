# FastAPI Production Tutorial — Complete Learning + Deployment Package

This repository turns the official FastAPI Tutorial topics into one structured,
production-oriented reference application. The examples are original implementations;
they are organized with separation of concerns instead of putting everything in `main.py`.

## Coverage

First steps, path/query parameters, request bodies, validation, nested Pydantic models,
cookies, headers, response models/status codes, forms/files, errors, dependencies and
yield dependencies, OAuth2/JWT/password hashing, middleware/CORS, SQLAlchemy/PostgreSQL,
multiple routers, JSON Lines streaming, SSE, background tasks, static files, templates,
WebSockets, testing, lifespan, settings, custom responses, OpenAPI metadata, Docker,
Jenkins, Kubernetes, AWS and Terraform.

## Architecture

```text
Internet -> CDN/WAF/ALB -> Kubernetes Ingress -> FastAPI
                                              |
                        +---------------------+--------------------+
                        |                     |                    |
                     Routers              Services           Dependencies
                        |                     |                    |
                     Schemas             Repositories          Security
                        |                     |                    |
                        +---------------------+--------------------+
                                              |
                                       PostgreSQL
```

## Separation of concerns

```text
app/api/routers      HTTP endpoints only
app/core             settings/security
app/db               engine/session/base
app/models           SQLAlchemy persistence models
app/schemas          Pydantic contracts
app/repositories     database access
app/services         business logic
app/dependencies     FastAPI dependency providers
app/middleware       cross-cutting HTTP concerns
app/workers          background work
```

## Local

```bash
uv sync
cp .env.example .env
docker compose up -d postgres
uv run fastapi dev app/main.py
```

Swagger: http://localhost:8000/docs

## Tests

```bash
uv run pytest -q
uv run ruff check .
```

## Docker

```bash
docker compose up --build
```

## Kubernetes

```bash
kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/secret.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/ingress.yaml
kubectl apply -f k8s/hpa.yaml
```

The sample Kubernetes secret is for learning only. Use AWS Secrets Manager,
External Secrets, or an equivalent secret-management system in production.

## AWS/Terraform

The Terraform folder provides a VPC foundation and documents the production path:
Route53 -> CloudFront/WAF -> ALB -> EKS -> FastAPI -> RDS PostgreSQL.
It deliberately leaves production-specific IAM, EKS and database choices explicit.

## Jenkins

The Jenkinsfile demonstrates checkout -> lint -> tests -> Docker build -> ECR push ->
EKS rollout. Configure AWS/Jenkins credentials securely; never commit access keys.

## Production rules

Keep business logic out of routers; use Pydantic at the boundary, services for business
workflows, repositories for persistence, dependency injection for shared/request-scoped
resources, migrations for schemas, structured logs/metrics/traces, readiness/liveness
probes, rate limits, TLS, secret management, non-root containers, resource limits and
safe retry/idempotency policies.
