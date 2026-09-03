PROTO=api/proto/order/v1/order.proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO)
tidy:
	go mod tidy
test:
	go test ./...
run:
	go run ./cmd/server
build:
	go build -o bin/order-service ./cmd/server
docker:
	docker build -t order-service:latest .
