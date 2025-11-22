# Next Steps Checklist

Use this checklist to finish hardening the starter after cloning it into an environment with module download access.

## Dependency hygiene
- Run `go mod tidy` to materialize `go.sum` and confirm module integrity.
- If your environment is restricted, set up a private module proxy or vendor dependencies with `go mod vendor` and commit the `vendor/` directory.

## Database workflow
- Update `config.yaml` (and environment variables) with your real PostgreSQL connection details.
- Create new Goose migrations for your domain using `make migrate-create name=add_feature` and apply them with `make migrate-up`.
- Add SQLC queries under `internal/db/queries` and regenerate code with `make sqlc`.

## Application wiring
- Replace the in-memory user repository with a SQL-backed implementation in `internal/repository` when ready.
- Extend domain services under `internal/domain` and keep handlers thin by delegating to them.
- Keep routes versioned under `internal/http/routes` and document any new endpoints in `internal/docs/openapi.yaml`.

## Quality gates
- Run `make test` and `make lint` once dependencies are available.
- Update or add handler and domain tests alongside new functionality.

## Documentation
- Refresh `README.md` with any project-specific setup notes.
- Regenerate Scalar docs with `make docs` after updating the OpenAPI spec.
