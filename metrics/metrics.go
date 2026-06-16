package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	OnlineUsersCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "threexui_total_online_users",
		Help: "Total number of online users",
	})
)
