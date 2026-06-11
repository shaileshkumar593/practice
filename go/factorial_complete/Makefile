run:
	go run factorial.go

test:
	go test ./...

build:
	go build -o app

docker-test:
	docker build -t factorial-test .
	docker run --rm factorial-test
