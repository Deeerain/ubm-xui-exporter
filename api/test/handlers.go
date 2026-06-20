package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Deeerain/ubm-xui-exporter/api"
)

func OnlinesHandler(mockToken string, users ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mockResponse := api.ApiResponse[[]string]{
			Obj: users,
		}

		if r.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", mockToken) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "Application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}
}

func UniqueIpsHandler(mockToken string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mockResponse := api.ApiResponse[[]api.ClientIpInfo]{
			Obj: []api.ClientIpInfo{
				{
					Id: 1,
					Ips: []api.IPInfo{
						{Ip: "192.168.0.1"},
						{Ip: "192.168.0.2"},
					},
				},
			},
		}

		if r.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", mockToken) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "Application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}
}

func ServerStatusHandler(mockToken string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mockResponse := api.ApiResponse[api.ServerStatus]{
			Obj: api.ServerStatus{
				Cpu: 1234,
				Mem: api.MemoryStatus{
					Current: 1234,
					Total:   1234,
				},
			},
		}

		if r.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", mockToken) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}
}

func CreateMockXuiServer(mockToken string, secretPath string) *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc(makeUrl(secretPath, "/panel/api/clients/onlines"), OnlinesHandler(mockToken, "test-1", "test-2"))
	mux.HandleFunc(makeUrl(secretPath, "/panel/api/server/clientIps"), UniqueIpsHandler(mockToken))
	mux.HandleFunc(makeUrl(secretPath, "/panel/api/server/status"), ServerStatusHandler(mockToken))

	return httptest.NewServer(mux)
}

func makeUrl(mockSecretPath, path string) string {
	return fmt.Sprintf("%s%s", mockSecretPath, path)
}
