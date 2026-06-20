package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"slices"
)

type APIClientOpts struct {
	AccessToken string
	BaseURL     string
	SecretPath  string
}

type APIClient struct {
	opts       APIClientOpts
	httpClient *http.Client
}

func NewAPIClient(opts APIClientOpts, httpClient *http.Client) *APIClient {
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}

	return &APIClient{
		opts:       opts,
		httpClient: httpClient,
	}
}

func (c *APIClient) GetOnlineUsersCount() (int, error) {
	body, err := c.doRequest("/panel/api/clients/onlines", http.MethodPost)
	if err != nil {
		return -1, fmt.Errorf("Failed to get online users: %w", err)
	}

	var bodyObj ApiResponse[[]string]
	if err := json.Unmarshal(body, &bodyObj); err != nil {
		return -1, fmt.Errorf("Failed to unmarshal response: %w", err)
	}

	return len(bodyObj.Obj), nil
}

func (c *APIClient) GetUniqueIps() ([]string, error) {
	body, err := c.doRequest("/panel/api/server/clientIps", http.MethodGet)
	if err != nil {
		return nil, fmt.Errorf("Failed to get client ips: %w", err)
	}
	var bodyObj ApiResponse[[]ClientIpInfo]

	if err := json.Unmarshal(body, &bodyObj); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response: %w", err)
	}

	var ips []string
	for _, clinet := range bodyObj.Obj {
		for _, ip := range clinet.Ips {
			if slices.Contains(ips, ip.Ip) {
				continue
			}

			ips = append(ips, ip.Ip)
		}
	}

	return ips, err
}

func (c *APIClient) GetServerStatus() (*ServerStatus, error) {
	body, err := c.doRequest("/panel/api/server/status", http.MethodGet)
	if err != nil {
		return nil, fmt.Errorf("Failed to get server status: %w", err)
	}

	var bodyObj ApiResponse[ServerStatus]

	if err := json.Unmarshal(body, &bodyObj); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal body: %w", err)
	}

	return &bodyObj.Obj, nil
}

func (c *APIClient) createRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.opts.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (c *APIClient) doRequest(path string, method string) ([]byte, error) {
	url := fmt.Sprintf("%s%s%s", c.opts.BaseURL, c.opts.SecretPath, path)
	req, err := c.createRequest(method, url)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Warn("Failed to get response", "status_code", resp.StatusCode)

		return nil, fmt.Errorf("Received non-OK response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	slog.Debug("Requst", "url", url, "status_code", resp.StatusCode)

	return body, nil
}
