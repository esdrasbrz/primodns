package config

import (
	"errors"
	"os"
)

type SonarrConfig struct {
	Addr   string
	ApiKey string
}

func loadSonarrConfig() (*SonarrConfig, error) {
	config := &SonarrConfig{
		Addr:   os.Getenv("SONARR_ADDR"),
		ApiKey: os.Getenv("SONARR_API_KEY"),
	}

	if config.Addr == "" || config.ApiKey == "" {
		return nil, errors.New("Missing Sonarr address or API key environment variables.")
	}

	return config, nil
}
