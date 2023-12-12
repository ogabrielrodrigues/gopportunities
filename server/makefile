.PHONY: default run migrate_up migrate_down

DATABASE_DRIVER="postgres"
DATABASE_URL=$(shell grep -m 1 database config.toml | tr -s ' ' | tr -d '"' | tr -d "'" | cut -d' ' -f3)
MIGRATIONS_PATH="config/database/${DATABASE_DRIVER}/migration"

default: run

run:
	@go run cmd/api/main.go

migrate_up:
	@migrate -path ${MIGRATIONS_PATH} -database "${DATABASE_URL}" -verbose up

migrate_down:
	@migrate -path ${MIGRATIONS_PATH} -database "${DATABASE_URL}" -verbose down
