package config

import (
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Cloudflare  *CloudflareConfig
	MetricsPort int
}

func LoadConfig() (*Config, error) {
	// Load envfile .env variables if it exists
	godotenv.Load()

	cloudflareConfig, err := loadCloudflareConfig()
	if err != nil {
		return nil, err
	}

	httpMetricsPort, err := strconv.Atoi(getEnv("HTTP_METRICS_PORT", "9987"))
	if err != nil {
		return nil, err
	}

	config := Config{
		Cloudflare:  cloudflareConfig,
		MetricsPort: httpMetricsPort,
	}

	return &config, nil
}
