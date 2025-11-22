# AI Coding Agent: Project Rules & Responsibilities

This document defines how an AI coding agent must operate when contributing to this repository. The goal is to keep the codebase consistent, reliable, and production-friendly.

---

## 1. Architectural Requirements

The agent must follow the existing **layered architecture**:

* `cmd/` – entrypoints only
* `internal/config` – configuration logic only
* `internal/http` – routing, handlers, middleware, response normalization
* `internal/domain` – business logic, use cases
* `internal/repository` – database access, SQLC code
* `pkg/` – shared packages (logger, errors)

Do not place business logic inside handlers or repository functions.

---

## 2. Conventions

### Routing

* All routes must be versioned (`/v1/...`)
* Group routes by domain
* Use Fiber idioms (`ctx.Params`, `ctx.BodyParser`, etc.)

### Error Handling

* All errors must use the unified structured format:

  ```json
  {"error": {"code": "ERR_CODE", "message": "...", "details": {}}}
  ```
* Never leak internal DB errors directly to the client.

### Logging

* Use the project's slog logger.
* Log errors at handler or domain layer, not repository layer.

### Configuration

* All new config fields must be added to:

  * config struct
  * config.yaml example
  * environment variable mapping

---

## 3. Database and Migrations

### SQLC

* Create SQL queries inside `internal/db/queries`
* Regenerate with: `make sqlc`
* Wrap SQLC methods in repository interfaces for testability

### Goose

* Each new domain must have its own migration file
* Naming:
  `YYYYMMDDHHMM__description.sql`

---

## 4. API Documentation (Scalar)

The agent must maintain API documentation:

* Every new endpoint must be documented
* Every request/response type must have an OpenAPI schema
* `make docs` must generate clean output with no missing references

---

## 5. Testing Rules

* Handlers must have basic request/response tests
* Repository tests should use SQLite in-memory or Docker Postgres
* Domain logic must include unit tests when applicable
* Avoid mocking SQLC directly; mock repository interfaces instead

---

## 6. Coding Style Expectations

* Prefer pure, stateless functions in domain layer
* Prefer small, focused files
* Avoid global state
* Follow Go naming conventions
* Use dependency injection (constructor functions)

---

## 7. Makefile Requirements

If the agent adds functionality requiring build steps, it must update the Makefile with new tasks when needed.

---

## 8. Allowed & Forbidden Actions

### Allowed

* Create new endpoints
* Extend domain logic
* Add migrations
* Update configuration
* Add tests
* Maintain documentation
* Add optional modules (auth, cache, queues)

### Forbidden

* Breaking directory structure
* Writing business logic inside handlers
* Returning raw internal errors to API consumers
* Adding libraries without updating README
* Changing config behavior without documentation

---

## 9. Pull Request Behavior

When generating or assisting with PRs, the agent must:

* Show only relevant diffs
* Explain reasoning briefly
* Validate that migrated SQL is reversible
* Ensure tests pass for all affected areas
* Ensure code compiles before submission

---

## 10. Golden Rules Summary

* Keep handlers thin
* Keep domain pure
* Keep repository simple
* Keep errors structured
* Keep logs structured
* Keep code maintainable
* Keep migrations clean
* Keep docs updated
