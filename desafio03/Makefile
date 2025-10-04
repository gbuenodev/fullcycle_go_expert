# Variáveis
DB_URL=mysql://root:root@tcp(localhost:3306)/orders?multiStatements=true

.PHONY: all up down run generate test migrate-create migrate-up migrate-down

all: up run

# Docker
up:
	docker-compose up -d

down:
	docker-compose down

# Go Application
run:
	cd cmd/ordersystem && go run main.go wire_gen.go

generate:
	go generate ./...

test:
	go test ./...

# Database Migrations
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

migrate-up: up
	@echo "Waiting for MySQL to be ready..."
	@while ! docker exec mysql mysqladmin ping -h"127.0.0.1" --silent; do \
			sleep 1; \
	done
	@echo "MySQL is ready. Running migrations..."
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down: up
	migrate -path migrations -database "$(DB_URL)" -verbose down