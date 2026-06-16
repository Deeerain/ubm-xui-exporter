package config

import "os"

type Config struct {
	ListenAddress string
	MetricsPath   string
}

func ParseEnv() (*Config, error) {
	config := Config{}

	config.ListenAddress = parseEnvOrDefault("LISTEN_ADDRESS", ":8080")
	config.MetricsPath = parseEnvOrDefault("METRICS_PATH", "/metrics")

	return &config, nil
}

func parseEnvOrDefault(envKey, defaultValue string) string {
	if value := os.Getenv(envKey); value != "" {
		return value
	}
	return defaultValue
}
