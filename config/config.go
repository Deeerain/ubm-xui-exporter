package config

import (
	"fmt"
	"os"
)

type Config struct {
	ListenAddress  string
	MetricsPath    string
	XUIBaseURL     string
	XUISecretPath  string
	XUIAccessToken string
}

func ParseEnv() (*Config, error) {
	config := Config{}

	config.ListenAddress = parseEnvOrDefault("LISTEN_ADDRESS", ":8080")
	config.MetricsPath = parseEnvOrDefault("METRICS_PATH", "/metrics")
	config.XUIBaseURL = parseEnvOrDefault("XUI_BASE_URL", "http://localhost:8080")
	config.XUISecretPath = parseEnvOrDefault("XUI_SECRET_PATH", "/")
	config.XUIAccessToken = parseEnvOrDefault("XUI_ACCESS_TOKEN", "")

	if config.XUIAccessToken == "" {
		return nil, fmt.Errorf("XUI_ACCESS_TOKEN is required")
	}

	return &config, nil
}

func parseEnvOrDefault(envKey, defaultValue string) string {
	if value := os.Getenv(envKey); value != "" {
		return value
	}
	return defaultValue
}
