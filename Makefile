MAKEFLAGS += --silent

.PHONY: run

run:
	DB_USERNAME=romanhuk \
	DB_PASSWORD=qwerty \
	DB_HOST=localhost  \
	DB_NAME=go_restish_dev \
	DB_PORT=5432 \
	go run cmd/server/main.go

dev/db/create:
	createdb go_restish_dev

test/db/create:
	createdb go_restish_test

dev/db/migrate:
	tern migrate -c ./internal/database/dev.conf -m ./internal/database/migrations

dev/db/rollback:
	tern migrate -c ./internal/database/dev.conf -m ./internal/database/migrations -d -1
