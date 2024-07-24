package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	Sonarr *SonarrConfig
	Radarr *RadarrConfig
}

func LoadConfig() (*Config, error) {
	// Load envfile .env variables if it exists
	godotenv.Load()

	sonarrConfig, err := loadSonarrConfig()
	if err != nil {
		return nil, err
	}

	radarrConfig, err := loadRadarrConfig()
	if err != nil {
		return nil, err
	}

	config := Config{
		Sonarr: sonarrConfig,
		Radarr: radarrConfig,
	}

	return &config, nil
}