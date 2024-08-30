package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	Cloudflare *CloudflareConfig
}

func LoadConfig() (*Config, error) {
	// Load envfile .env variables if it exists
	godotenv.Load()

	cloudflareConfig, err := loadCloudflareConfig()
	if err != nil {
		return nil, err
	}

	config := Config{
		Cloudflare: cloudflareConfig,
	}

	return &config, nil
}
