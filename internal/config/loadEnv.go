package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}

func LoadDBConfig() (string, string, string, string, string) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "admin"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "321"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "db_login"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	return host, user, password, dbname, port
}
