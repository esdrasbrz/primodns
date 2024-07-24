package config

import (
	"errors"
	"os"
)

type RadarrConfig struct {
	Addr string
	ApiKey string
}

func loadRadarrConfig() (*RadarrConfig, error) {
	config := &RadarrConfig{
		Addr: os.Getenv("RADARR_ADDR"),
		ApiKey: os.Getenv("RADARR_API_KEY"),
	}

	if config.Addr == "" || config.ApiKey == "" {
		return nil, errors.New("Missing Radarr address or API key environment variables.")
	}

	return config, nil
}