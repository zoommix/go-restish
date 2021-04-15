MAKEFLAGS += --silent

.PHONY: run

run:
	DB_USERNAME=romanhuk \
	DB_PASSWORD=qwerty \
	DB_HOST=localhost  \
	DB_NAME=fasthttp_template_development \
	DB_PORT=5432 \
	go run cmd/server/main.go
