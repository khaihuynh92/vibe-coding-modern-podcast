package config

import (
	"os"
	"strings"
)

// Config holds all configuration for the application
type Config struct {
	Port        string
	Environment string
	CORSOrigins []string
	LogLevel    string
}

// Load loads configuration from environment variables with sensible defaults
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "3001"),
		Environment: getEnv("GO_ENV", "development"),
		CORSOrigins: getCORSOrigins(),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getCORSOrigins parses CORS origins from environment variable
func getCORSOrigins() []string {
	origins := getEnv("CORS_ORIGINS", "http://localhost:3000")
	return strings.Split(origins, ",")
}

// IsDevelopment returns true if running in development mode
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsProduction returns true if running in production mode
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}
