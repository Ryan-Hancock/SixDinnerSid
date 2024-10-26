.PHONY: up down reset migrate-up migrate-down

# Docker commands
up:
	docker-compose up -d

down:
	docker-compose down

reset: down
	docker-compose down -v
	docker-compose up -d

# Database migrations
migrate-up:
	migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/cat_feeder?sslmode=disable" up

migrate-down:
	migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/cat_feeder?sslmode=disable" down

# Create a new migration file
create-migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

# Development helpers
dev: up migrate-up
	go run main.go

test:
	go test -v ./...