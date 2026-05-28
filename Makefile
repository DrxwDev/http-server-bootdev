include .env

DB_DSN=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

migrations:
	@goose -dir=internal/migrations/schema create ${name} sql

mu: migrate/up
migrate/up:
	@echo "Running up migrations..."
	@goose -dir=internal/migrations/schema postgres ${DB_DSN} up

md: migrate/down
migrate/down:
	@echo "Rolling back migrations..."
	@goose -dir=internal/migrations/schema postgres ${DB_DSN} down

mf: migrate/fresh
migrate/fresh:
	@echo "Dropping..."
	@goose -dir=internal/migrations/schema postgres ${DB_DSN} reset
	@echo "Migrating..."
	@$(MAKE) --no-print-directory migrate/up
