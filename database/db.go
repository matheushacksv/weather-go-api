package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect creates a connection with the database
func Connect(database_url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), database_url)
	if err != nil {
		return nil, errors.New("unable to connect to the database")
	}
	return pool, nil
}
