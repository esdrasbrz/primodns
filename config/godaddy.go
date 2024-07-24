package config

import (
	"errors"
	"os"
)

type GodaddyConfig struct {
	ApiKey string
	ApiSecret string
	TargetDomain string
}

func loadGodaddyConfig() (*GodaddyConfig, error) {
	config := &GodaddyConfig{
		ApiKey: os.Getenv("GODADDY_API_KEY"),
		ApiSecret: os.Getenv("GODADDY_API_SECRET"),
		TargetDomain: os.Getenv("GODADDY_TARGET_DOMAIN"),
	}

	if config.ApiKey == "" || config.ApiSecret == "" || config.TargetDomain == "" {
		return nil, errors.New("Missing GoDaddy environment variables.")
	}

	return config, nil
}