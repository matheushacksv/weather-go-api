-include .env
export

DB_URL ?= postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
TEST_DB_URL = postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/weather_test?sslmode=disable

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migrate-test-up:
	migrate -path database/migrations -database "$(TEST_DB_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" down

migrate-force:
	migrate -path database/migrations -database "$(DB_URL)" force $(VERSION)

test:
	TEST_DATABASE_URL="$(TEST_DB_URL)" go test -v ./api/users
