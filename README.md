# Go API Starter

Production-ready Fiber starter kit with Viper configuration, slog logging, SQLC-ready database layer, and Scalar docs.

## Features
- Layered architecture (`cmd`, `internal`, `pkg`).
- Structured error responses and slog-based request logging.
- Fiber middleware for CORS, request ID, rate limiting, and panic recovery.
- Health and user example endpoints with OpenAPI documentation.
- SQLC query definitions and Goose migrations for PostgreSQL.
- Docker Compose for API + Postgres and hot reload via Air.

## Getting Started
### Prerequisites
- Go 1.22+
- make
- (optional) Docker + docker-compose

### Local Run
```bash
make run
```
The server listens on the port configured in `config.yaml` (default `8080`).

### Hot Reload
```bash
make dev
```
Requires [Air](https://github.com/cosmtrek/air).

### Docker Compose
```bash
docker-compose up --build
```
This starts the API and PostgreSQL. Override configuration via environment variables prefixed with `APP_` (e.g., `APP_DB_HOST`).

## Configuration
Configuration is loaded from `config.yaml` and environment variables (`APP_` prefix). Key fields:
- `server.port` – HTTP port
- `db.*` – database connection details
- `logging.level` – debug|info|warn|error
- `features.*` – toggle optional modules (auth, cache, queues)

## Database Workflow
- Define migrations in `internal/db/migrations` using Goose naming (`YYYYMMDDHHMM__description.sql`).
- Create SQLC queries in `internal/db/queries` and run `make sqlc` to generate type-safe code into `internal/repository/sqlc`.
- Migration commands:
  - `make migrate-up`
  - `make migrate-down`
  - `make migrate-create name=add_feature`

## API
- Health check: `GET /v1/health`
- Users CRUD under `/v1/users`
- Docs served from `internal/docs/openapi.yaml` at `/docs`.

Error responses follow:
```json
{"error": {"code": "ERR_CODE", "message": "...", "details": {}}}
```

## Testing & Linting
- `make test`
- `make lint` (requires `golangci-lint`)

## SQLC & Scalar
- `make sqlc` generates repository code from queries.
- `make docs` can be wired to a Scalar CLI to regenerate the OpenAPI bundle.
