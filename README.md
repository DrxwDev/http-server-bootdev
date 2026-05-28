# Chirpy

A Go-based HTTP API server for a Twitter-like microblogging service, built as part of the [Boot.dev](https://boot.dev) curriculum.

## Tech Stack

- **Go** — Core language
- **Gin** — HTTP router
- **PostgreSQL** — Database
- **pgx** — PostgreSQL driver with connection pooling
- **sqlc** — SQL code generation
- **goose** — Database migrations
- **uber-go/fx** — Dependency injection
- **air** — Live-reload during development

## Getting Started

### Prerequisites

- Go 1.26+
- PostgreSQL
- [air](https://github.com/air-verse/air) (for live-reload)
- [goose](https://github.com/pressly/goose) (for migrations)
- [sqlc](https://sqlc.dev) (for code generation)

### Setup

1. Clone the repository and enter the directory.
2. Copy `.env.example` to `.env` (if it exists) or configure the environment variables listed in `.env`:

   ```
   HOST=localhost
   PORT=8080
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=chirpy
   PLATFORM=dev
   ```

3. Create the database:
   ```bash
   createdb chirpy
   ```

4. Run database migrations:
   ```bash
   make migrate/up
   ```

5. Start the server:
   ```bash
   go run ./cmd/api/
   ```
   Or with live-reload:
   ```bash
   air
   ```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/app/` | Welcome page (tracked by hit counter middleware) |
| GET | `/app/assets/logo.png` | Static asset |
| POST | `/api/v1/users` | Create a new user |
| POST | `/api/validate_chirp` | Validate and profanity-filter a chirp message |
| POST | `/admin/reset` | Reset the users table (dev only) |

## Project Structure

```
cmd/api/         — Application entrypoint
internal/
├── bootdev/     — Basic endpoints (health, chirp validation)
├── config/      — Configuration loading
├── database/    — sqlc-generated code and connection pool
├── logger/      — Structured logging
├── middlewares/ — HTTP middlewares (metrics, hit counting, reset)
├── migrations/  — Database schema and SQL queries
├── server/      — Gin router setup and route registration
└── users/       — Users module (controller → service → repository)
assets/          — Static files (HTML, images)
```

## Available Commands

```bash
make migrate/up      # Apply pending migrations
make migrate/down    # Roll back last migration
make migrate/fresh   # Drop and re-run all migrations
make migrations      # Create a new migration (name=<name>)
sqlc generate        # Regenerate Go code from SQL queries
```

## Profanity Filter

The `/api/validate_chirp` endpoint replaces the words `kerfuffle`, `sharbert`, and `fornax` with `****`.
