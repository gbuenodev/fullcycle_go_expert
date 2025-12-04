.PHONY: build run-local docker-build docker-run test

build:
	go build -o bin/loadtest ./cmd/loadtest

run-local:
	go run ./cmd/loadtest --url=http://google.com --requests=10 --concurrency=2

docker-build:
	docker build -t loadtest:latest .

docker-run:
	docker run --rm loadtest:latest \
		--url=http://google.com \
		--requests=100 \
		--concurrency=10

test:
	go test -v -race ./...