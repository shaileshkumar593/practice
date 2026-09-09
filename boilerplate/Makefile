APP=go-kit-boilerplate

.PHONY: run test build migrate-up migrate-down fmt vet tidy

run:
	go run ./cmd/api

test:
	go test ./... -cover

build:
	go build -o bin/api ./cmd/api

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/appdb?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/appdb?sslmode=disable" down 1
