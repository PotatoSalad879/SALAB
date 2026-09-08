package config

import (
	"fmt"
	"os"
)

// Config holds application settings loaded from environment variables.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBTimeZone string
	ServerPort string
}

// LoadConfig reads environment variables and returns a typed configuration.
func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("SERVER_PORT")
	}
	if port == "" {
		port = "8000"
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "learnhub"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "learnhub"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "learnhub"
	}

	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	timeZone := os.Getenv("DB_TIMEZONE")
	if timeZone == "" {
		timeZone = "Asia/Bangkok"
	}

	return &Config{
		DBHost:     host,
		DBPort:     dbPort,
		DBUser:     user,
		DBPassword: password,
		DBName:     dbname,
		DBSSLMode:  sslMode,
		DBTimeZone: timeZone,
		ServerPort: port,
	}, nil
}

// DSN returns a PostgreSQL connection string derived from the Config.
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort, c.DBSSLMode, c.DBTimeZone,
	)
}
