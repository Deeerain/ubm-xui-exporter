package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Deeerain/ubm-xui-exporter/api"
	"github.com/Deeerain/ubm-xui-exporter/api/test"
	"github.com/Deeerain/ubm-xui-exporter/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	mockToken      string = "test-token"
	mockSecretPath string = "/secret"
	exporterServer *httptest.Server
)

func TestMain(m *testing.M) {
	xuiServer := test.CreateMockXuiServer(mockToken, mockSecretPath)
	defer xuiServer.Close()

	mockClient := api.NewAPIClient(api.APIClientOpts{
		SecretPath:  mockSecretPath,
		AccessToken: mockToken,
		BaseURL:     xuiServer.URL,
	}, nil)

	mux := http.NewServeMux()
	reg := prometheus.NewRegistry()
	collector := metrics.NewXUICollector(mockClient)
	reg.MustRegister(collector)
	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	exporterServer = httptest.NewServer(mux)

	code := m.Run()

	exporterServer.Close()
	xuiServer.Close()
	os.Exit(code)
}

func Test_Metrics(t *testing.T) {
	resp, err := http.Get(exporterServer.URL + "/metrics")
	if err != nil {
		t.Fatal("Failed to get response", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Not 200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Cant read body: %v", body)
	}

	t.Logf("Metrics:\n%s", string(body))
}
