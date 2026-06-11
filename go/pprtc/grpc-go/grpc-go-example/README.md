# Go gRPC Example

This project demonstrates:

- gRPC server
- gRPC client
- Protocol Buffers
- Unary RPC
- Simple production-ready folder structure

## Install protoc plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Generate protobuf code

```bash
protoc --go_out=. --go-grpc_out=. proto/hello.proto
```

## Run server

```bash
go run server/main.go
```

## Run client

```bash
go run client/main.go
```
