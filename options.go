package aiozai

import (
	"net/http"
	"time"
)

// ClientConfig holds the resolved configuration for the SDK client.
type ClientConfig struct {
	APIKey        string
	BaseURL       string
	Timeout       time.Duration
	UploadTimeout time.Duration
	RetryConfig   *RetryConfig
	HTTPClient    *http.Client
}

// RetryConfig controls the retry behavior of the SDK client.
type RetryConfig struct {
	MaxRetries           int
	BaseDelay            time.Duration
	MaxDelay             time.Duration
	RetryableStatusCodes []int
}

// Option is a functional option for configuring the SDK client.
type Option func(*ClientConfig)

// WithAPIKey sets the API key for authenticated endpoints.
func WithAPIKey(apiKey string) Option {
	return func(c *ClientConfig) {
		c.APIKey = apiKey
	}
}

// WithBaseURL overrides the default API base URL.
func WithBaseURL(baseURL string) Option {
	return func(c *ClientConfig) {
		c.BaseURL = baseURL
	}
}

// WithTimeout sets the request timeout for standard API calls.
func WithTimeout(timeout time.Duration) Option {
	return func(c *ClientConfig) {
		c.Timeout = timeout
	}
}

// WithUploadTimeout sets the timeout for upload operations.
func WithUploadTimeout(timeout time.Duration) Option {
	return func(c *ClientConfig) {
		c.UploadTimeout = timeout
	}
}

// WithRetryConfig overrides the default retry configuration.
func WithRetryConfig(rc *RetryConfig) Option {
	return func(c *ClientConfig) {
		c.RetryConfig = rc
	}
}

// WithHTTPClient sets a custom HTTP client for the SDK.
func WithHTTPClient(client *http.Client) Option {
	return func(c *ClientConfig) {
		c.HTTPClient = client
	}
}

// defaultConfig returns the default client configuration.
func defaultConfig() *ClientConfig {
	return &ClientConfig{
		BaseURL:       DefaultBaseURL,
		Timeout:       DefaultTimeout,
		UploadTimeout: DefaultUploadTimeout,
		RetryConfig: &RetryConfig{
			MaxRetries:           DefaultMaxRetries,
			BaseDelay:            DefaultBaseDelay,
			MaxDelay:             DefaultMaxDelay,
			RetryableStatusCodes: RetryableStatusCodes,
		},
	}
}
