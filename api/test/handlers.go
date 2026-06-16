package test

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func OnlinesHandler(mockToken string, users ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mockResponse struct {
			Obj []string
		}

		mockResponse.Obj = []string(users)

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
		type IpInfo struct {
			Ip string `json:"ip"`
		}
		type ClientInfo struct {
			Id  int      `json:"id"`
			Ips []IpInfo `json:"ips"`
		}
		var mockResponse struct {
			Obj []ClientInfo `json:"obj"`
		}

		mockResponse.Obj = []ClientInfo{
			{
				Id: 1,
				Ips: []IpInfo{
					{Ip: "192.168.0.1"},
					{Ip: "192.168.0.2"},
				},
			},
			{
				Id: 2,
				Ips: []IpInfo{
					{Ip: "192.168.0.1"},
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
