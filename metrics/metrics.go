package metrics

import (
	"log/slog"
	"time"

	"github.com/Deeerain/ubm-xui-exporter/api"
	"github.com/prometheus/client_golang/prometheus"
)

type XUICollector struct {
	apiClient     *api.APIClient
	onlineDesc    *prometheus.Desc
	uniqueIpsDesc *prometheus.Desc
	cpuUsageDesc  *prometheus.Desc
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
		cpuUsageDesc: prometheus.NewDesc(
			"threexui_cpu_usage",
			"",
			nil, nil,
		),
		apiClient: apiClient,
	}
}

func (c *XUICollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.onlineDesc
	ch <- c.uniqueIpsDesc
	ch <- c.cpuUsageDesc
}

func (c *XUICollector) Collect(ch chan<- prometheus.Metric) {
	startTime := time.Now()

	slog.Debug("Starting scrap metrics")

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

	serverStatus, err := c.apiClient.GetServerStatus()
	if err != nil {
		slog.Error("Failed to get metrics", "error", err)
		return
	}

	slog.Debug("Status", "value", serverStatus)

	ch <- prometheus.MustNewConstMetric(
		c.cpuUsageDesc,
		prometheus.GaugeValue,
		float64(serverStatus.Cpu),
	)

	slog.Info("Successfuly scraped metrics", "duration_ms", time.Since(startTime).Microseconds())
}
