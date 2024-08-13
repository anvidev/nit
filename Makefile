include .env

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
export GOOSE_MIGRATION_DIR=migrations

db-status:
	@goose status

db-up:
	@goose up

db-down:
	@goose down

db-val:
	@goose validate

db-create:
	@goose create $(name) sql

db-reset:
	@goose reset
