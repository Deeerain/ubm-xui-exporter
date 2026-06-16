package metrics

import (
	"log/slog"

	"github.com/Deeerain/ubm-xui-exporter/api"
	"github.com/prometheus/client_golang/prometheus"
)

type XUICollector struct {
	apiClient     *api.APIClient
	onlineDesc    *prometheus.Desc
	uniqueIpsDesc *prometheus.Desc
}

func NewXUICollector(apiClient *api.APIClient) *XUICollector {
	return &XUICollector{
		onlineDesc: prometheus.NewDesc(
			"threexui_online_user_count",
			"Currnet number of online users connected to x-ui.",
			nil, nil,
		),
		uniqueIpsDesc: prometheus.NewDesc(
			"threexui_unique_ip_count",
			"Current member of unique ips connected to x-ui.",
			nil, nil,
		),
		apiClient: apiClient,
	}
}

func (c *XUICollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.onlineDesc
	ch <- c.uniqueIpsDesc
}

func (c *XUICollector) Collect(ch chan<- prometheus.Metric) {
	currentUsers, err := c.apiClient.GetOnlineUsersCount()
	if err != nil {
		slog.Error("Failed to get metrics", "error", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.onlineDesc,
		prometheus.GaugeValue,
		float64(currentUsers),
	)

	uniqueIps, err := c.apiClient.GetUniqueIps()
	if err != nil {
		slog.Error("Failed to get metrics", "error", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.uniqueIpsDesc,
		prometheus.GaugeValue,
		float64(len(uniqueIps)),
	)
}
