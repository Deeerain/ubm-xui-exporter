package main

import (
	"log/slog"
	"net/http"

	"github.com/Deeerain/ubm-xui-exporter/api"
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
	apiClient := api.NewAPIClient(api.APIClientOpts{
		BaseURL:     config.XUIBaseURL,
		SecretPath:  config.XUISecretPath,
		AccessToken: config.XUIAccessToken,
	}, nil)
	mux := http.NewServeMux()
	registry := prometheus.NewRegistry()

	collector := metrics.NewXUICollector(apiClient)
	registry.MustRegister(
		collector,
	)

	mux.Handle(config.MetricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	if err := http.ListenAndServe(config.ListenAddress, mux); err != nil {
		slog.Error("Failed to serve", "error", err)
	}
}
