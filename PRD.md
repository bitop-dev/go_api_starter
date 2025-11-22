# **Product Requirements Document (PRD)

Golang API Starter Template**

---

## **1. Overview**

This project provides a production-ready, opinionated Golang API starter template optimized for **rapid API prototyping** that can scale into a **production service** with minimal refactoring. It includes:

* Golang
* Fiber (HTTP framework)
* SQLC (type-safe DB layer)
* Goose (database migrations)
* Scalar (OpenAPI documentation)
* PostgreSQL as the primary database
* Dockerized development environment

The template is designed to be cloned and used as a starting point for building small-to-medium backend services.

---

## **2. Goals**

* Accelerate early API prototyping.
* Provide a clean, layered structure that transitions smoothly into production.
* Enforce consistent standards (structure, tooling, error format, logging, API design).
* Enable fast iteration with a repeatable and predictable development workflow.
* Allow optional modules (auth, caching, queues) to be added later without breaking the base structure.

---

## **3. Non-Goals / Out of Scope**

* No CI/CD configuration.
* No project bootstrap script beyond Makefile tasks.
* No advanced distributed-system features (service mesh, multi-service orchestration).
* No autoscaling or infrastructure provisioning.

---

## **4. Target Users**

* Engineers building new APIs quickly.
* Teams designing backend prototypes intended to grow into production.
* Developers looking for a well-structured template with modern tooling already configured.

---

# **5. Functional Requirements**

## **5.1 Project Structure**

### **Layered Architecture**

```
/cmd/api            # entrypoint
/internal/
    config/         # Viper config loader
    http/           
        middleware/ # logging, recover, cors, rate limits
        routes/     # route grouping, versioning
        handlers/   # Fiber handlers
        responses/  # standardized API response helpers
    domain/         # business logic
    repository/     # SQLC wrappers + repository interfaces
    db/
        migrations/ # Goose migrations
        queries/    # SQLC queries
    docs/           # Scalar/OpenAPI config
/pkg/
    logger/         # slog setup
    errors/         # structured error responses
```

---

## **5.2 API Structure**

### **Middleware (preconfigured)**

* Request logging (slog)
* Panic recovery
* CORS
* Rate limiting
* Request ID injection
* Structured error normalization

### **Routing**

* Versioned routing:
  `/v1/users`, `/v1/health`, …
* Routes separated by domain in `/internal/http/routes`

### **Scalar API Documentation**

* OpenAPI 3 generation automated on `make docs`
* Example endpoints included
* Hosted at `/docs`

---

## **5.3 Database**

### **Database Support**

**Primary:** PostgreSQL
**Optional:** SQLite example configs

### **SQLC Configuration**

* Example schema and queries included
* Generated code placed in `/internal/repository/sqlc`
* Repository interfaces wrapping SQLC for testability

### **Goose Migrations**

* Baseline migration included
* Naming conventions: `YYYYMMDDHHMM__description.sql`
* Makefile tasks:

  * `make migrate-up`
  * `make migrate-down`
  * `make migrate-create name=add_users`

---

## **5.4 Configuration System**

### **Config Management**

* Viper-based
* Reads from:

  * `config.yaml`
  * environment variables
* Loads:

  * app settings
  * DB config
  * logging options
  * server port
  * feature flags

Example:

```yaml
server:
  port: 8080

db:
  driver: postgres
  host: localhost
  port: 5432
  user: app
  pass: password
  name: appdb

logging:
  level: info
```

---

## **5.5 Logging**

### **Logging Framework**

* Uses **slog**
* Configurable log level via config
* Request logs (method, path, duration, status)
* Structured JSON logs

---

## **5.6 Error Handling**

### **Error Behavior**

* All handlers return a unified structured response:

```json
{
  "error": {
    "code": "USER_NOT_FOUND",
    "message": "User does not exist",
    "details": {}
  }
}
```

* Internal errors are logged, not leaked.

---

## **5.7 Example Functionality Included**

### Included baseline features:

* `/health` endpoint
* `/docs` endpoint
* Example user module:

  * migrations for `users`
  * SQLC queries
  * repository implementation
  * basic CRUD handlers
  * documented in Scalar

---

# **6. Optional Modules (included but disabled)**

Feature flags in config:

* `auth.jwt.enable`
* `cache.redis.enable`
* `queue.nats.enable`

Each includes:

* Folder structure
* Placeholder interfaces
* Example config
* No active code unless enabled

---

# **7. Development Workflow**

## **Local Development (Docker)**

* `docker-compose up` runs:

  * Go API (using air for live reload)
  * Postgres
* Hot reload support
* DB migrations applied on startup (optional flag)

## **Makefile Tasks**

```
make run          # local run without Docker
make dev          # run with hot reload
make build
make test
make lint
make sqlc
make migrate-up
make migrate-down
make docs         # generate Scalar docs
```

---

# **8. Quality Gates**

## **Linting**

* golangci-lint preconfigured
* Recommended rules:

  * govet
  * errcheck
  * gofumpt
  * revive
  * staticcheck

## **Testing**

* Baseline `testing` package
* Simple mock repository example
* Example handler tests

---

# **9. Non-Functional Requirements**

* Clean file layout
* Predictable repository interfaces
* Safe error-handling approach
* Logging structured for production
* API responses consistent across all endpoints
* Simple enough for prototyping, structured enough for production

---

# **10. Deliverables**

This PRD mandates the following deliverables in the final repo:

### **Documentation**

* `README.md`
* API usage guide
* DB migration guide
* SQLC overview
* Config reference
* Example workflows

### **Code**

* Fully structured project folder
* Sample endpoints
* Baseline database schema
* Makefile
* Docker setup
* Scalar docs
* Example tests
* Linter config

### **Config**

* `config.yaml` default
* environment variable mapping

### **Migrations**

* Base schema + examples included

---

If you want, I can now also generate:

* A full **README.md**
* The recommended **directory tree**
* A **scaffold generator plan**
* The **first set of baseline endpoints and migrations** (users, health, etc.)

Tell me what you want next.
