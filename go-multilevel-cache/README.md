# Go Multi-Level Cache Example

Production-oriented learning example showing:

- L1 in-process local cache
- L2 Redis distributed cache
- Cache-aside pattern
- TTL
- Bounded local cache
- Request coalescing with `singleflight`
- Cache invalidation
- FastAPI-like HTTP endpoints implemented with Go `net/http`
- Docker Compose Redis

## Architecture

Client
  |
  v
Go API
  |
  +--> L1 Local Memory Cache
  |
  +--> L2 Redis
  |
  +--> Repository / Database

Read path:

L1 HIT -> return

L1 MISS
  -> Redis HIT
  -> populate L1
  -> return

L1 MISS
  -> Redis MISS
  -> singleflight
  -> repository/database
  -> Redis SET
  -> L1 SET
  -> return

Write path:

Database update
  -> L1 delete
  -> Redis delete

## Prerequisites

- Go 1.24+
- Docker / Docker Compose

## Run Redis

```bash
docker compose up -d
```

## Run the application

```bash
go mod tidy
go run .
```

The API starts at:

http://localhost:8080

## Test

Health:

```bash
curl http://localhost:8080/health
```

First request:

```bash
curl http://localhost:8080/users/101
```

Expected source:

```json
"source": "DB"
```

Second request:

```bash
curl http://localhost:8080/users/101
```

Expected source:

```json
"source": "L1"
```

To test L2, wait 30 seconds for L1 TTL to expire and call again. Redis should still contain the value:

```json
"source": "L2"
```

## Update and invalidate

```bash
curl -X POST http://localhost:8080/users/update \
  -H "Content-Type: application/json" \
  -d '{"id":101,"name":"Alice Updated","email":"alice.updated@example.com"}'
```

The update modifies the repository and deletes both L1 and L2 cache entries.

## Important production considerations

This repository uses an in-memory map for the database and a simple bounded eviction policy for L1 so the project stays self-contained.

For production:

1. Replace the repository with PostgreSQL.
2. Use a real LRU/TinyLFU cache for L1.
3. Add cache metrics: hit/miss, latency, evictions, Redis errors.
4. Add distributed L1 invalidation using Kafka/Pub/Sub when multiple pods are running.
5. Add TTL jitter to reduce synchronized expirations.
6. Consider stale-while-revalidate/background refresh for hot keys.
7. Add Redis Sentinel/Cluster or a managed Redis service for high availability.
8. Never treat Redis as the source of truth.
9. Add timeouts and circuit-breaking around Redis.
10. Use structured logging and OpenTelemetry tracing.

## Key interview concepts

L1:
- Fastest
- Process-local
- Not shared between pods
- Very low latency

L2:
- Shared
- Redis
- Network hop
- Shared between application instances

Request coalescing:
- Prevents multiple identical concurrent cache misses from hitting the database simultaneously.
- `singleflight` is process-local. Cross-pod coordination requires a distributed mechanism.

Cache invalidation:
- Database is updated first.
- Cache entries are invalidated.
- In multi-pod systems, propagate invalidation events so every L1 cache is cleared.
