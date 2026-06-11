# Order gRPC Service with gRPC-Gateway and Swagger

Production-oriented starter project that implements:
- gRPC service for Orders (Get, Create, Update)
- In-memory slice storage (thread-safe)
- gRPC-Gateway to expose RESTful HTTP
- OpenAPI (swagger) YAML example
- `Makefile` for generation steps

## Requirements
- Go 1.20+
- protoc (protocol buffers compiler)
- protoc plugins:
  - protoc-gen-go
  - protoc-gen-go-grpc
  - protoc-gen-grpc-gateway
  - protoc-gen-openapiv2 (for swagger generation)

Install (example):
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

## Generate code
From project root:
```bash
make proto
```

This calls `protoc` to generate:
- `proto/order.pb.go`
- `proto/order_grpc.pb.go`
- `proto/order.pb.gw.go`
- `proto/order.swagger.json` (OpenAPI)

## Run
First generate code, then:
```bash
go run ./cmd/server
```

gRPC: `localhost:50051`
HTTP gateway: `localhost:8080` (REST endpoints)
Swagger/OpenAPI file: `proto/order.swagger.json` (generated)

## Notes
- This project uses an in-memory slice protected by a mutex for demo/testing.
- For production replace store with persistent DB, add structured logging, metrics, tracing, auth.
