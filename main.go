package main

import (
	"log/slog"
	"net/http"

	"github.com/Deeerain/ubm-xui-exporter/config"
	"github.com/Deeerain/ubm-xui-exporter/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config, err := config.ParseEnv()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	registry := prometheus.NewRegistry()
	registry.MustRegister(
		metrics.OnlineUsersCount,
	)

	mux.Handle(config.MetricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	if err := http.ListenAndServe(config.ListenAddress, mux); err != nil {
		slog.Error("Failed to serve", "error", err)
	}
}
