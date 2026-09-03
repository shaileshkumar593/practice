# Production gRPC Order Service

Go gRPC reference project with protobuf, PostgreSQL, clean architecture,
interceptors, graceful shutdown, Docker and Kubernetes-ready design.

## Generate protobuf
Install protoc, protoc-gen-go and protoc-gen-go-grpc, then run:
make proto

## Run
Set DATABASE_URL and run:
make tidy
make run

## Architecture
Client -> LB -> gRPC interceptors -> Handler -> Application Service
-> Repository -> PostgreSQL.

Redis caching, OAuth/OIDC authentication, TLS/mTLS, OpenTelemetry,
Prometheus, idempotency persistence, outbox/Kafka and full streaming
implementation should be added according to production requirements.
