package test

import (
	"encoding/json"
	"fmt"
	"net/http"

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
