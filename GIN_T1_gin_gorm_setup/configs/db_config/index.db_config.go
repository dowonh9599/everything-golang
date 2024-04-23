package db_config

import "os"

var (
	DB_HOST     = "localhost"
	DB_PORT     = "5432" // Default port for PostgreSQL
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test_db"
)

func InitDatabaseConfig() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost != "" {
		DB_HOST = dbHost
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort != "" {
		DB_PORT = os.Getenv("DB_PORT")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser != "" {
		DB_USER = dbUser
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword != "" {
		DB_PASSWORD = dbPassword
	}

	dbName := os.Getenv("DB_NAME")
	if dbName != "" {
		DB_NAME = dbName
	}
}
