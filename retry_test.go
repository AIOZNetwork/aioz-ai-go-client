package aiozai

import (
	"context"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockTransport is a test RoundTripper that returns configurable responses.
type mockTransport struct {
	responses []*http.Response
	errors    []error
	callCount atomic.Int32
}

func (m *mockTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	idx := int(m.callCount.Add(1)) - 1
	if idx < len(m.errors) && m.errors[idx] != nil {
		return nil, m.errors[idx]
	}
	if idx < len(m.responses) {
		return m.responses[idx], nil
	}
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(""))}, nil
}

func TestRetryTransport_NoRetryOn200(t *testing.T) {
	mock := &mockTransport{
		responses: []*http.Response{
			{StatusCode: 200, Body: io.NopCloser(strings.NewReader("ok"))},
		},
	}

	rt := newRetryTransport(mock, &RetryConfig{
		MaxRetries:           3,
		BaseDelay:            10 * time.Millisecond,
		MaxDelay:             100 * time.Millisecond,
		RetryableStatusCodes: RetryableStatusCodes,
	})

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com/test", nil)
	resp, err := rt.RoundTrip(req)

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(1), mock.callCount.Load())
}

func TestRetryTransport_NoRetryOn400(t *testing.T) {
	mock := &mockTransport{
		responses: []*http.Response{
			{StatusCode: 400, Body: io.NopCloser(strings.NewReader("bad request"))},
		},
	}

	rt := newRetryTransport(mock, &RetryConfig{
		MaxRetries:           3,
		BaseDelay:            10 * time.Millisecond,
		MaxDelay:             100 * time.Millisecond,
		RetryableStatusCodes: RetryableStatusCodes,
	})

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com/test", nil)
	resp, err := rt.RoundTrip(req)

	require.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, int32(1), mock.callCount.Load())
}

func TestRetryTransport_RetriesOn500(t *testing.T) {
	mock := &mockTransport{
		responses: []*http.Response{
			{StatusCode: 500, Body: io.NopCloser(strings.NewReader("error"))},
			{StatusCode: 500, Body: io.NopCloser(strings.NewReader("error"))},
			{StatusCode: 200, Body: io.NopCloser(strings.NewReader("ok"))},
		},
	}

	rt := newRetryTransport(mock, &RetryConfig{
		MaxRetries:           3,
		BaseDelay:            10 * time.Millisecond,
		MaxDelay:             100 * time.Millisecond,
		RetryableStatusCodes: RetryableStatusCodes,
	})

	req, _ := http.NewRequestWithContext(context.Background(), "POST", "http://example.com/test", strings.NewReader("body"))
	resp, err := rt.RoundTrip(req)

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(3), mock.callCount.Load())
}

func TestRetryTransport_RetriesOn502(t *testing.T) {
	mock := &mockTransport{
		responses: []*http.Response{
			{StatusCode: 502, Body: io.NopCloser(strings.NewReader("bad gateway"))},
			{StatusCode: 200, Body: io.NopCloser(strings.NewReader("ok"))},
		},
	}

	rt := newRetryTransport(mock, &RetryConfig{
		MaxRetries:           2,
		BaseDelay:            10 * time.Millisecond,
		MaxDelay:             100 * time.Millisecond,
		RetryableStatusCodes: RetryableStatusCodes,
	})

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com/test", nil)
	resp, err := rt.RoundTrip(req)

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(2), mock.callCount.Load())
}

func TestRetryTransport_MaxRetriesExhausted(t *testing.T) {
	mock := &mockTransport{
		responses: []*http.Response{
			{StatusCode: 503, Body: io.NopCloser(strings.NewReader("unavailable"))},
			{StatusCode: 503, Body: io.NopCloser(strings.NewReader("unavailable"))},
			{StatusCode: 503, Body: io.NopCloser(strings.NewReader("unavailable"))},
			{StatusCode: 503, Body: io.NopCloser(strings.NewReader("unavailable"))},
		},
	}

	rt := newRetryTransport(mock, &RetryConfig{
		MaxRetries:           3,
		BaseDelay:            10 * time.Millisecond,
		MaxDelay:             100 * time.Millisecond,
		RetryableStatusCodes: RetryableStatusCodes,
	})

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com/test", nil)
	resp, err := rt.RoundTrip(req)

	require.NoError(t, err)
	assert.Equal(t, 503, resp.StatusCode)
	assert.Equal(t, int32(4), mock.callCount.Load()) // 1 initial + 3 retries
}

func TestRetryTransport_ContextCancellation(t *testing.T) {
	mock := &mockTransport{
		responses: []*http.Response{
			{StatusCode: 500, Body: io.NopCloser(strings.NewReader("error"))},
			{StatusCode: 500, Body: io.NopCloser(strings.NewReader("error"))},
		},
	}

	rt := newRetryTransport(mock, &RetryConfig{
		MaxRetries:           5,
		BaseDelay:            1 * time.Second, // Long delay to ensure context cancels first
		MaxDelay:             10 * time.Second,
		RetryableStatusCodes: RetryableStatusCodes,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", "http://example.com/test", nil)
	_, err := rt.RoundTrip(req)

	assert.Error(t, err)
	assert.True(t, mock.callCount.Load() <= 2, "should stop retrying after context cancellation")
}

func TestRetryTransport_NetworkError(t *testing.T) {
	mock := &mockTransport{
		errors: []error{
			assert.AnError,
			nil,
		},
		responses: []*http.Response{
			nil,
			{StatusCode: 200, Body: io.NopCloser(strings.NewReader("ok"))},
		},
	}

	rt := newRetryTransport(mock, &RetryConfig{
		MaxRetries:           3,
		BaseDelay:            10 * time.Millisecond,
		MaxDelay:             100 * time.Millisecond,
		RetryableStatusCodes: RetryableStatusCodes,
	})

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com/test", nil)
	resp, err := rt.RoundTrip(req)

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, int32(2), mock.callCount.Load())
}

func TestRetryTransport_NilConfigNoRetry(t *testing.T) {
	mock := &mockTransport{
		responses: []*http.Response{
			{StatusCode: 500, Body: io.NopCloser(strings.NewReader("error"))},
		},
	}

	rt := newRetryTransport(mock, nil)

	req, _ := http.NewRequestWithContext(context.Background(), "GET", "http://example.com/test", nil)
	resp, err := rt.RoundTrip(req)

	require.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)
	assert.Equal(t, int32(1), mock.callCount.Load())
}

func TestComputeDelay(t *testing.T) {
	baseDelay := 100 * time.Millisecond
	maxDelay := 5 * time.Second

	// Attempt 0: ~100-125ms
	d0 := computeDelay(0, baseDelay, maxDelay)
	assert.GreaterOrEqual(t, d0, baseDelay)
	assert.LessOrEqual(t, d0, 125*time.Millisecond)

	// Attempt 1: ~200-250ms
	d1 := computeDelay(1, baseDelay, maxDelay)
	assert.GreaterOrEqual(t, d1, 200*time.Millisecond)
	assert.LessOrEqual(t, d1, 250*time.Millisecond)

	// Attempt 2: ~400-500ms
	d2 := computeDelay(2, baseDelay, maxDelay)
	assert.GreaterOrEqual(t, d2, 400*time.Millisecond)
	assert.LessOrEqual(t, d2, 500*time.Millisecond)
}

func TestComputeDelay_CapsAtMaxDelay(t *testing.T) {
	baseDelay := 1 * time.Second
	maxDelay := 5 * time.Second

	// Large attempt should cap at maxDelay
	d := computeDelay(20, baseDelay, maxDelay)
	assert.LessOrEqual(t, d, maxDelay)
}

func TestIsRetryableStatus(t *testing.T) {
	assert.True(t, isRetryableStatus(500, RetryableStatusCodes))
	assert.True(t, isRetryableStatus(502, RetryableStatusCodes))
	assert.True(t, isRetryableStatus(503, RetryableStatusCodes))
	assert.True(t, isRetryableStatus(504, RetryableStatusCodes))
	assert.False(t, isRetryableStatus(200, RetryableStatusCodes))
	assert.False(t, isRetryableStatus(400, RetryableStatusCodes))
	assert.False(t, isRetryableStatus(401, RetryableStatusCodes))
	assert.False(t, isRetryableStatus(404, RetryableStatusCodes))
}
