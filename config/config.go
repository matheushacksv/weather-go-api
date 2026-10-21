package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config represents the project config
type Config struct {
	Port     string
	DBString string
}

// Load gets the data from env and returns a Config
func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, errors.New("erro to read .env")
	}

	port := os.Getenv("PORT")
	DBName := os.Getenv("POSTGRES_DB")
	DBUser := os.Getenv("POSTGRES_USER")
	DBPassword := os.Getenv("POSTGRES_PASSWORD")
	DBHost := os.Getenv("POSTGRES_HOST")
	DBPort := os.Getenv("POSTGRES_PORT")

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
		DBUser, DBPassword, DBHost, DBPort, DBName,
	)

	if port == "" {
		port = "8080"
	}

	return Config{
		port,
		databaseURL,
	}, nil
}
