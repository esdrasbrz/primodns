package config

import (
	"errors"
	"os"
	"strings"
)

type CloudflareConfig struct {
	ApiToken   string
	ZoneId     string
	DnsRecords []string
}

func loadCloudflareConfig() (*CloudflareConfig, error) {
	config := &CloudflareConfig{
		ApiToken:   os.Getenv("CLOUDFLARE_API_TOKEN"),
		ZoneId:     os.Getenv("CLOUDFLARE_ZONE_ID"),
		DnsRecords: strings.Split(os.Getenv("CLOUDFLARE_DNS_RECORDS"), ","),
	}

	if config.ApiToken == "" || config.ZoneId == "" || len(config.DnsRecords) == 0 {
		return nil, errors.New("Missing Cloudflare environment variables.")
	}

	return config, nil
}
