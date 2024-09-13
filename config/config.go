package config

import (
	"fmt"
	"os"
)

// Config holds the database configuration
type Config struct {
	DB_Host     string
	DB_User     string
	DB_Password string
	DB_Name     string
	DB_Port     string
	SSLMode     string
}

// LoadConfig reads configuration from environment variables
func LoadConfig() (*Config, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, fmt.Errorf("DB_HOST environment variable is required")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, fmt.Errorf("DB_USER environment variable is required")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD environment variable is required")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("DB_NAME environment variable is required")
	}

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		return nil, fmt.Errorf("DB_PORT environment variable is required")
	}
	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "" {
		return nil, fmt.Errorf("SSL_MODE environment variable is required")
	}

	config := &Config{
		DB_Host:     dbHost,
		DB_User:     dbUser,
		DB_Password: dbPassword,
		DB_Name:     dbName,
		DB_Port:     dbPortStr,
		SSLMode:     sslMode, // Optionally process this further if needed
	}

	return config, nil
}
