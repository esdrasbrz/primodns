package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	LastUpdatedAt = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "primodns_last_updated_at",
		Help: "The timestamp when the IP address was updated",
	}, []string{"ip"})
	ExternalIPRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "primodns_request_external_ip_total",
		Help: "The number of requests to external IP service",
	}, []string{"status"})
	CloudflareRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "primodns_request_cloudflare_total",
		Help: "The number of requests to Cloudflare",
	}, []string{"code", "dns_record", "ip"})
)

func ServeMetrics(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(addr, nil)
}
