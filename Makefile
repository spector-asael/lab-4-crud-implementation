# CMPS3162 Advanced Databases - Banking System
# Asael Tobar
# February 23rd, 2025

# Pass in the .envrc file, which exports BANK_DB_DSN
include .envrc

## run: run the cmd/api application
.PHONY: run help checkbalance deposit comment healthcheck all
run: 
	@echo 'Running application...'
	@go run ./cmd/api

# Help target
help:
	@echo ""
	@echo "Application:"
	@echo "  make run            - Run API server"
	@echo ""
	@echo "API Testing (requires server running):"
	@echo "  make checkbalance   - Test POST /v1/balance (user 1)"
	@echo "  make checkbalance2  - Test POST /v1/balance (user 2)"
	@echo "  make deposit        - Test POST /v1/deposit (valid request)"
	@echo "  make deposit2       - Test POST /v1/deposit (invalid request)"
	@echo "  make comment        - Test POST /v1/comments"
	@echo "  make healthcheck    - Test GET  /v1/healthcheck"
	@echo "  make all            - Run all API tests"
	@echo ""
	@echo "Database:"
	@echo "  make db/psql        - Connect to database using psql"
	@echo ""
	@echo "Migrations:"
	@echo "  make db/migrations/new name=NAME"
	@echo "                      - Create new migration"
	@echo ""
	@echo "  make migrations/up  - Apply all migrations"
	@echo "  make migrations/down"
	@echo "                      - Revert all migrations"
	@echo ""
	@echo "  make migrations/fix version=VERSION"
	@echo "                      - Force schema_migrations to VERSION"
	@echo ""
	@echo "  make migrations/init"
	@echo "                      - Generate all initial banking tables"
	@echo ""
	@echo "Environment:"
	@echo "  Requires BANK_DB_DSN exported in .envrc"
	@echo ""

# POST /balance
checkbalance:
	@echo "Testing /balance..."
	curl -X POST http://localhost:4000/v1/balance \
	-H "Content-Type: application/json" \
	-d '{"user_id":1,"bank_number":111111}'

checkbalance2:
	@echo "Testing /balance..."
	curl -X POST http://localhost:4000/v1/balance \
	-H "Content-Type: application/json" \
	-d '{"user_id":2,"bank_number":111111}'

# POST /deposit
deposit:
	@echo "Testing /deposit..."
	curl -X POST http://localhost:4000/v1/deposit \
	-H "Content-Type: application/json" \
	-d '{"user_id":1,"bank_number":111111,"deposit_amount":500.75}'

deposit2:
	@echo "Testing /deposit..."
	curl -X POST http://localhost:4000/v1/deposit \
	-H "Content-Type: application/json" \
	-d '{"user_id":1,"bank_number":111111,"deposit_amount":500 75}'

# POST /comments
comment:
	@echo "Testing /comments..."
	curl -X POST http://localhost:4000/v1/comments \
	-H "Content-Type: application/json" \
	-d '{"content":"This is a test comment","author":"Spector"}'

# GET /healthcheck
healthcheck:
	@echo "Testing /healthcheck..."
	curl -X GET http://localhost:4000/v1/healthcheck

## db/psql: Connect to the banking database using psql
.PHONY: db
db/psql:
	psql ${BANK_DB_DSN}

## db/migrations/new name=$1: Create a new database migration
.PHONY: migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## migrations/up: Apply all up database migrations
.PHONY: migrations/up
migrations/up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${BANK_DB_DSN} up

## db/migrations/down: Revert all migrations
.PHONY: migrations/down
migrations/down:
	@echo 'Reverting all migrations...'
	migrate -path ./migrations -database ${BANK_DB_DSN} down

## db/migrations/fix version=$1: Force schema_migrations version
.PHONY: migrations/fix
migrations/fix:
	@echo 'Forcing schema migrations version to ${version}...'
	migrate -path ./migrations -database ${BANK_DB_DSN} force ${version}