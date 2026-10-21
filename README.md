![CI](https://github.com/matheushacksv/weather_api/actions/workflows/ci.yml/badge.svg)
# Weather API Go with users

## Description
An API with third party API for weather in Go for studying of Gin

### Stack
- Go
- Gin
- Postgres
- pgx v5
- migrate
- docker compose
- Make

### Configuration

Env file:
```env
# Postgres
POSTGRES_DB=postgres
POSTGRES_USER=postgres_user
POSTGRES_PASSWORD=postgres
POSTGRES_HOST=localhost
POSTGRES_PORT=5432

# Server
PORT=3002
```

### How to run
1. Clone the repo
2. Copy .env
3. Run postgres `docker compose up -d`
4. Migrate the database using Make:
  `make migrate-up`
5. Run the api `go run .`

### Endpoints

| Method | Route | Description |
|--------|-------|-------------|
| GET | `/health` | API and database health check |
| POST | `/users` | Create a user |
| GET | `/users` | List all users |
| GET | `/users/:ID` | Get a user by ID |
| PUT | `/users/:ID` | Update a user |
| DELETE | `/users/:ID` | Delete a user |
| GET | `/weather?latitude={lat}&longitude={long}` | Current weather from Open-Meteo |
