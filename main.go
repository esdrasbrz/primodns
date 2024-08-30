package main

import (
	"github.com/esdrasbrz/primodns/config"
	"github.com/esdrasbrz/primodns/ddns"
	"github.com/esdrasbrz/primodns/services/cloudflare"
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

	cfs := cloudflare.New(config.Cloudflare, logger)
	dyn := ddns.New(logger, cfs)

	dyn.RunDDNSUpdater()
}
