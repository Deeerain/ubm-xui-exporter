package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Deeerain/ubm-xui-exporter/api"
)

var (
	mockToken      string = "test-token"
	mockSecretPath string = "/secret"
	client         *api.APIClient
)

func TestMain(m *testing.M) {
	mux := http.NewServeMux()

	mux.HandleFunc(makeUrl("/panel/api/clients/onlines"), OnlinesHandler(mockToken, "test-1", "test-2"))
	mux.HandleFunc(makeUrl("/panel/api/server/clientIps"), UniqueIpsHandler(mockToken))

	server := httptest.NewServer(mux)
	defer server.Close()

	client = api.NewAPIClient(api.APIClientOpts{
		AccessToken: mockToken,
		BaseURL:     server.URL,
		SecretPath:  mockSecretPath,
	}, nil)

	m.Run()
}

func makeUrl(path string) string {
	return fmt.Sprintf("%s%s", mockSecretPath, path)
}

func Test_GetOnlineUsers(t *testing.T) {
	count, err := client.GetOnlineUsersCount()

	if err != nil {
		t.Fatal(err)
	}

	if count != 2 {
		t.Fatalf("Count error")
	}
}

func Test_GetUniqueIps(t *testing.T) {
	ips, err := client.GetUniqueIps()

	if err != nil {
		t.Fatal(err)
	}

	if len(ips) == 0 {
		t.Error("Not unique ips")
	}
}
