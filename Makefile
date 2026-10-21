-include .env
export

DB_URL ?= postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" down
