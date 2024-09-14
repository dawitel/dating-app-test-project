package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)
// Config represests the cofigurations vars for the app.
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int
	SSLMode    string
	Port       string
	JWTSecret  string
}

// LoadConfig assembles the config and returns a pointer to it.
func LoadConfig() *Config {
	port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "1234"),
		DBName:     getEnv("DB_NAME", "dating_app"),
		DBPort:     port,
		SSLMode:    getEnv("SSL_MODE", "disabled"),
		Port:       getEnv("PORT", ":8080"),
		JWTSecret:  getEnv("JWT_SECRET", "topsecret"),
	}
}

// getEnv returs the values of the env variales.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetDSN returns a GORM-compatible DSN string.
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort, c.SSLMode,
	)
}
