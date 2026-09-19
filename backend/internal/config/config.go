package config

import (
	"os"

	env "github.com/joho/godotenv"
)

type Config struct {
	// Add configuration fields here, e.g., database connection strings, API keys, etc.
	HTTPPort string `json:"http_port"`
}

func LoadConfig() (*Config, error) {
	// Load configuration from environment variables, files, or other sources.
	// For example, you can read from a JSON or YAML file:
	_ = env.Load(".env.local", ".env")

	// Parse the configuration file and populate the Config struct.
	var config Config
	config.HTTPPort = getEnv("HTTP_PORT", "8080")

	return &config, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
