package main

import (
	"net/http"

	"github.com/Deeerain/ubm-xui-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config, err := config.ParseEnv()
	if err != nil {
		panic(err)
	}
	registry := prometheus.NewRegistry()
	mux := http.NewServeMux()

	mux.Handle(config.MetricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.ListenAndServe(config.ListenAddress, mux)
}
