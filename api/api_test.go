package api

import (
	"testing"

	"github.com/Deeerain/ubm-xui-exporter/config"
)

var client *APIClient

func init() {
	cfg, err := config.ParseEnv()
	if err != nil {
		panic(err)
	}

	client = NewAPIClient(APIClientOpts{
		AccessToken: cfg.XUIAccessToken,
		BaseURL:     cfg.XUIBaseURL,
		SecretPath:  cfg.XUISecretPath,
	}, nil)
}

func Test_GetOnlineUsers(t *testing.T) {
	if err := client.GetOnlineUsersCount(); err != nil {
		t.Fatal(err)
	}
}
