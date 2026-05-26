package aiozai

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient_DefaultConfig(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err)

	cfg := client.Config()
	assert.Equal(t, DefaultBaseURL, cfg.BaseURL)
	assert.Equal(t, DefaultTimeout, cfg.Timeout)
	assert.Equal(t, DefaultUploadTimeout, cfg.UploadTimeout)
	assert.Equal(t, DefaultMaxRetries, cfg.RetryConfig.MaxRetries)
	assert.Equal(t, DefaultBaseDelay, cfg.RetryConfig.BaseDelay)
	assert.Equal(t, DefaultMaxDelay, cfg.RetryConfig.MaxDelay)
}

func TestNewClient_WithAPIKey(t *testing.T) {
	client, err := NewClient(WithAPIKey("test-api-key"))
	require.NoError(t, err)

	cfg := client.Config()
	assert.Equal(t, "test-api-key", cfg.APIKey)
}

func TestNewClient_WithBaseURL(t *testing.T) {
	customURL := "https://custom.api.com/api/v1"
	client, err := NewClient(WithBaseURL(customURL))
	require.NoError(t, err)

	cfg := client.Config()
	assert.Equal(t, customURL, cfg.BaseURL)
}

func TestNewClient_WithTimeout(t *testing.T) {
	client, err := NewClient(WithTimeout(60 * time.Second))
	require.NoError(t, err)

	cfg := client.Config()
	assert.Equal(t, 60*time.Second, cfg.Timeout)
}

func TestNewClient_WithRetryConfig(t *testing.T) {
	rc := &RetryConfig{
		MaxRetries: 5,
		BaseDelay:  2 * time.Second,
		MaxDelay:   60 * time.Second,
	}
	client, err := NewClient(WithRetryConfig(rc))
	require.NoError(t, err)

	cfg := client.Config()
	assert.Equal(t, 5, cfg.RetryConfig.MaxRetries)
	assert.Equal(t, 2*time.Second, cfg.RetryConfig.BaseDelay)
	assert.Equal(t, 60*time.Second, cfg.RetryConfig.MaxDelay)
}

func TestNewClient_WithHTTPClient(t *testing.T) {
	customClient := &http.Client{Timeout: 120 * time.Second}
	client, err := NewClient(WithHTTPClient(customClient))
	require.NoError(t, err)

	cfg := client.Config()
	assert.NotNil(t, cfg.HTTPClient)
}

func TestNewClient_Raw(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err)

	raw := client.Raw()
	assert.NotNil(t, raw)
}

func TestNewClient_ServiceAccessors(t *testing.T) {
	client, err := NewClient()
	require.NoError(t, err)

	assert.NotNil(t, client.Models())
	assert.NotNil(t, client.Datasets())
	assert.NotNil(t, client.Competitions())
	assert.NotNil(t, client.Collections())
	assert.NotNil(t, client.Discussions())
	assert.NotNil(t, client.Notifications())
	assert.NotNil(t, client.Organizations())
	assert.NotNil(t, client.Repositories())
	assert.NotNil(t, client.Storage())
	assert.NotNil(t, client.Users())
	assert.NotNil(t, client.Public())
	assert.NotNil(t, client.Core())
}

func TestNewClient_InitPerformance(t *testing.T) {
	start := time.Now()
	_, err := NewClient(WithAPIKey("perf-test-key"))
	elapsed := time.Since(start)

	require.NoError(t, err)
	assert.Less(t, elapsed, 100*time.Millisecond, "Client init must complete in under 100ms")
}
