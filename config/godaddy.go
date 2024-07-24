package config

import (
	"errors"
	"os"
)

type GodaddyConfig struct {
	ApiKey string
	ApiSecret string
}

func loadGodaddyConfig() (*GodaddyConfig, error) {
	config := &GodaddyConfig{
		ApiKey: os.Getenv("GODADDY_API_KEY"),
		ApiSecret: os.Getenv("GODADDY_API_SECRET"),
	}

	if config.ApiKey == "" || config.ApiSecret == "" {
		return nil, errors.New("Missing GoDaddy API Key or Secret environment variables.")
	}

	return config, nil
}