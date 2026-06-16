package api

import (
	"fmt"
	"io"
	"net/http"
)

type APIClientOpts struct {
	AccessToken string
}

type APIClient struct {
	opts       APIClientOpts
	httpClient *http.Client
}

func NewAPIClient(opts APIClientOpts, httpClient *http.Client) *APIClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &APIClient{
		opts:       opts,
		httpClient: httpClient,
	}
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

func (c *APIClient) doRequest(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received non-OK response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	return body, nil
}
