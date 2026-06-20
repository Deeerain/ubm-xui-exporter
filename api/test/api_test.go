package test

import (
	"os"
	"testing"

	"github.com/Deeerain/ubm-xui-exporter/api"
)

var (
	mockToken      string = "test-token"
	mockSecretPath string = "/secret"
	client         *api.APIClient
)

func TestMain(m *testing.M) {
	server := CreateMockXuiServer(mockToken, mockSecretPath)

	client = api.NewAPIClient(api.APIClientOpts{
		AccessToken: mockToken,
		BaseURL:     server.URL,
		SecretPath:  mockSecretPath,
	}, nil)

	code := m.Run()

	server.Close()

	os.Exit(code)
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

func Test_GetServerStatus(t *testing.T) {
	status, err := client.GetServerStatus()

	if err != nil || status == nil {
		t.Fatal(err)
	}
}
