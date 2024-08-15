package cloudflare

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/esdrasbrz/primoflix/config"
	"github.com/hashicorp/go-multierror"
	"go.uber.org/zap"
)

type CloudflareService interface {
	UpdateDomains(content string) error
}

type service struct {
	config *config.CloudflareConfig
	logger *zap.Logger
}

func New(config *config.CloudflareConfig, logger *zap.Logger) CloudflareService {
	return &service{
		config: config,
		logger: logger,
	}
}

func (s *service) UpdateDomains(content string) (result error) {
	for _, dnsRecord := range s.config.DnsRecords {
		err := s.updateDomain(dnsRecord, content)

		if err != nil {
			result = multierror.Append(result, err)
		}
	}

	return
}

func (s *service) updateDomain(dnsRecord, content string) error {
	client := http.Client{}
	url := fmt.Sprintf(
		"https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s",
		s.config.ZoneId,
		dnsRecord,
	)
	body := []byte(fmt.Sprintf(`{"content": "%s"}`, content))

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + s.config.ApiToken},
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		s.logger.Error(
			"Error while updating domain",
			zap.String("dns_record", dnsRecord),
			zap.String("content", content),
			zap.Int("status_code", response.StatusCode),
		)

		return errors.New("Cloudflare update domain error")
	}

	return nil
}
