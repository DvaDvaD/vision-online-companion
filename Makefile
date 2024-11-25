include .env
export

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

.PHONY: run/db
run/db:
	@echo 'Running PostgreSQL container...'
	docker run -d --name my-postgres-container -p 5432:5432 my-postgres-image

.PHONY: stop/db
stop/db:
	@echo 'Stopping and removing PostgreSQL container...'
	docker stop my-postgres-container
	docker rm my-postgres-container

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path=./database/migrations -database=${DB_DSN} up

.PHONY: db/migrations/down
db/migrations/down:
	@echo 'Running down migrations...'
	migrate -path=./database/migrations -database=${DB_DSN} down

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./database/migrations ${name}

## db/seed: seed the database with initial data
.PHONY: db/seed
db/seed: confirm
	@echo 'Seeding the database...'
	psql -d ${DB_DSN} -f ./database/sql/seed/seed_1.sql
	psql -d ${DB_DSN} -f ./database/sql/seed/seed_2.sql

.PHONY: sqlc-gen
sqlc-gen:
	@echo 'Generating database code...'
	sqlc generate

.PHONY: oapi-gen
oapi-gen:
	@echo 'Generating OpenAPI spec compliant code...'
	oapi-codegen -package=gen -generate "types,spec,echo" docs/openapi.yml > api/gen/api.gen.go

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy and vendor dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...