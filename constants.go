package aiozai

import "time"

const (
	// Version is the current SDK version.
	Version = "1.0.4"

	// DefaultBaseURL is the default base URL for the AIOZ AI API.
	DefaultBaseURL = "https://api.aiozai.network/api/v1"

	// DefaultTimeout is the default request timeout.
	DefaultTimeout = 30 * time.Second

	// DefaultUploadTimeout is the default timeout for upload operations.
	DefaultUploadTimeout = 300 * time.Second

	// DefaultMaxRetries is the default number of retry attempts.
	DefaultMaxRetries = 3

	// DefaultBaseDelay is the initial delay for exponential backoff.
	DefaultBaseDelay = 1 * time.Second

	// DefaultMaxDelay is the maximum delay between retries.
	DefaultMaxDelay = 30 * time.Second
)

// RetryableStatusCodes contains HTTP status codes that trigger a retry.
var RetryableStatusCodes = []int{500, 502, 503, 504}
