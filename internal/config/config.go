package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	Port     string
	Host     string
	LogLevel string
}

// Load loads configuration from environment variables with defaults
func Load() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		Host:     getEnv("HOST", "0.0.0.0"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
