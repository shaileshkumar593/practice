Notes about proto generation:
- The `google/api/annotations.proto` file is required for HTTP options
  (it is provided by the googleapis repository).
- When using the `Makefile`, ensure `protoc` discovery path includes
  the protobuf include files. For example, if you installed googleapis
  locally:

  protoc -I proto -I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@*/third_party/googleapis ...

Adjust paths as needed.
