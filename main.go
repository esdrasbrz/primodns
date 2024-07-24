package main

import (
	"github.com/esdrasbrz/primoflix/config"
	"github.com/esdrasbrz/primoflix/services/cloudflare"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	config, err := config.LoadConfig()
	if err != nil {
		sugar.Fatal(err)
	}

	cloudflareService := cloudflare.New(config.Cloudflare, logger)

	err = cloudflareService.UpdateDomains("8.8.8.8")
	if err != nil {
		sugar.Fatal(err)
	}
}
