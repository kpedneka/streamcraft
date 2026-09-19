package config

import "testing"

func TestLoadConfig(t *testing.T) {
	// Test loading configuration from environment variables.
	config, err := LoadConfig()
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	if config.HTTPPort != "8080" {
		t.Errorf("expected HTTPPort to be '8080', got '%s'", config.HTTPPort)
	}
}
