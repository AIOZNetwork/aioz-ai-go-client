package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	aiozai "github.com/AIOZNetwork/aioz-ai-go-client"
)

func TestAuthFlow_APIKeyConfigured(t *testing.T) {
	client, err := aiozai.NewClient(
		aiozai.WithAPIKey("test-api-key"),
	)
	require.NoError(t, err)

	cfg := client.Config()
	assert.Equal(t, "test-api-key", cfg.APIKey)
	assert.NotNil(t, client.Raw())
}

func TestAuthFlow_NoAPIKeyConfigured(t *testing.T) {
	client, err := aiozai.NewClient()
	require.NoError(t, err)

	cfg := client.Config()
	assert.Empty(t, cfg.APIKey)
}

func TestClientCreation_WithCustomBaseURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client, err := aiozai.NewClient(
		aiozai.WithAPIKey("test-key"),
		aiozai.WithBaseURL(server.URL+"/api/v1"),
	)
	require.NoError(t, err)

	cfg := client.Config()
	assert.Equal(t, server.URL+"/api/v1", cfg.BaseURL)
}

func TestClientCreation_WithRetryConfig(t *testing.T) {
	client, err := aiozai.NewClient(
		aiozai.WithRetryConfig(&aiozai.RetryConfig{
			MaxRetries: 5,
			BaseDelay:  2 * time.Second,
			MaxDelay:   60 * time.Second,
		}),
	)
	require.NoError(t, err)

	cfg := client.Config()
	assert.Equal(t, 5, cfg.RetryConfig.MaxRetries)
	assert.Equal(t, 2*time.Second, cfg.RetryConfig.BaseDelay)
}
