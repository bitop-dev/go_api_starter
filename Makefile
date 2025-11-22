APP_NAME=go-api-starter

.PHONY: run dev build test lint sqlc migrate-up migrate-down migrate-create docs

run:
go run ./cmd/api

dev:
air

build:
go build -o bin/$(APP_NAME) ./cmd/api

test:
go test ./...

lint:
golangci-lint run ./...

sqlc:
sqlc generate

migrate-up:
goose -dir internal/db/migrations postgres "$${DATABASE_URL}" up

migrate-down:
goose -dir internal/db/migrations postgres "$${DATABASE_URL}" down

migrate-create:
goose -dir internal/db/migrations create "$${name}" sql

docs:
scalar generate internal/docs/openapi.yaml --output internal/docs
