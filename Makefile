MAKEFLAGS += --silent

.PHONY: run

run:
	go run cmd/server/main.go
