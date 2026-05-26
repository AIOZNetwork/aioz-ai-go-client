package aiozai

import (
	"bytes"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
	"slices"
	"time"
)

// retryTransport wraps an http.RoundTripper with retry logic using
// exponential backoff and jitter.
type retryTransport struct {
	next   http.RoundTripper
	config *RetryConfig
}

// newRetryTransport creates a new retry transport wrapping the given transport.
func newRetryTransport(next http.RoundTripper, config *RetryConfig) http.RoundTripper {
	if config == nil || config.MaxRetries <= 0 {
		return next
	}
	return &retryTransport{
		next:   next,
		config: config,
	}
}

// RoundTrip implements http.RoundTripper with retry logic.
func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body.Close()
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	}

	var lastResp *http.Response
	var lastErr error

	for attempt := range t.config.MaxRetries + 1 {
		if attempt > 0 {
			delay := computeDelay(attempt-1, t.config.BaseDelay, t.config.MaxDelay)

			select {
			case <-req.Context().Done():
				return nil, req.Context().Err()
			case <-time.After(delay):
			}

			// Reset body for retry
			if bodyBytes != nil {
				req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			}
		}

		resp, err := t.next.RoundTrip(req)
		if err != nil {
			lastErr = err
			lastResp = nil
			continue
		}

		if !isRetryableStatus(resp.StatusCode, t.config.RetryableStatusCodes) {
			return resp, nil
		}

		// Drain and close body before retry
		if resp.Body != nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}

		lastResp = resp
		lastErr = nil
	}

	if lastErr != nil {
		return nil, lastErr
	}
	return lastResp, nil
}

// computeDelay calculates the backoff delay with jitter for a given attempt.
// Formula: min(baseDelay * 2^attempt + jitter, maxDelay)
// Jitter: random 0-25% of computed delay.
func computeDelay(attempt int, baseDelay, maxDelay time.Duration) time.Duration {
	delay := float64(baseDelay) * math.Pow(2, float64(attempt))
	if delay > float64(maxDelay) {
		delay = float64(maxDelay)
	}

	// Add jitter: 0-25% of delay
	jitter := delay * 0.25 * rand.Float64()
	delay += jitter

	if time.Duration(delay) > maxDelay {
		return maxDelay
	}
	return time.Duration(delay)
}

// isRetryableStatus checks if the given status code is in the retryable list.
func isRetryableStatus(statusCode int, retryable []int) bool {
	return slices.Contains(retryable, statusCode)
}
