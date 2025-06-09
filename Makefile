# Makefile for running database migrations
# Assumes golang-migrate/migrate CLI is installed and .env file is present

# Load environment variables from .env file
include .env
export

# Construct PostgreSQL connection string (handle empty DB_PASSWORD)
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

# Default target
.PHONY: all
all: migrate

# Run database migrations
.PHONY: migrate
migrate:
	@echo "Running migrations..."
	@echo "Current directory: $(CURDIR)"
	@echo "Migration path: $(CURDIR)/db"
	@if [ ! -d "db" ]; then echo "Error: db directory not found"; exit 1; fi
	@if [ ! -f "db/000001_create_flood_tables.up.sql" ]; then echo "Error: migration file db/000001_create_flood_tables.up.sql not found"; exit 1; fi
	@ls -l $(CURDIR)/db
	@migrate -path $(CURDIR)/db -database "$(DB_URL)" -verbose up

#.PHONY: reset
#reset:
#	@echo "Resetting database..."
#	@if [ ! -d "db" ]; then echo "Error: db directory not found"; exit 1; fi
#	@if [ ! -f "db/000001_create_flood_tables.down.sql" ]; then echo "Error: migration file db/000001_create_flood_tables.down.sql not found"; exit 1; fi
#	@migrate -path $(CURDIR)/db -database "$(DB_URL)" -verbose down -all
