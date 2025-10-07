# Variáveis
DB_URL=mysql://root:root@tcp(localhost:3306)/orders?multiStatements=true

.PHONY: run docker-up docker-down server-up generate test migrate-create migrate-up migrate-down proto

run: docker-up server-up

# Docker
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Go Application
server-up:
	cd cmd/ordersystem && go run main.go wire_gen.go

generate:
	go generate ./...

test:
	go test ./...

# Database Migrations
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

migrate-up: docker-up
	@echo "Waiting for MySQL to be ready..."
	@while ! docker exec mysql mysqladmin ping -h"127.0.0.1" --silent; do \
			sleep 1; \
	done
	@echo "MySQL is ready. Running migrations..."
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down: docker-up
	migrate -path migrations -database "$(DB_URL)" -verbose down

# gRPC proto
proto:
	protoc --proto_path=internal/infra/grpc/protofiles --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/*.proto