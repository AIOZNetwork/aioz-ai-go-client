package aiozai

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAiozAPIError_Error(t *testing.T) {
	err := &AiozAPIError{
		StatusCode: 404,
		ErrorCode:  "NOT_FOUND",
		Message:    "Model not found",
		Endpoint:   "/api-key/model/123",
		Method:     "GET",
	}

	expected := "[404] NOT_FOUND: Model not found (GET /api-key/model/123)"
	assert.Equal(t, expected, err.Error())
}

func TestAiozAPIError_ErrorWithRequestID(t *testing.T) {
	err := &AiozAPIError{
		StatusCode: 500,
		ErrorCode:  "INTERNAL_ERROR",
		Message:    "Something went wrong",
		Endpoint:   "/api-key/model/list",
		Method:     "POST",
		RequestID:  "req-abc-123",
	}

	expected := "[500] INTERNAL_ERROR: Something went wrong (POST /api-key/model/list, request-id: req-abc-123)"
	assert.Equal(t, expected, err.Error())
}

func TestAiozAPIError_ImplementsError(t *testing.T) {
	var err error = &AiozAPIError{
		StatusCode: 401,
		ErrorCode:  "UNAUTHORIZED",
		Message:    "Invalid API key",
		Endpoint:   "/api-key/model/list",
		Method:     "POST",
	}

	assert.Error(t, err)
}

func TestAiozAPIError_ErrorsAs(t *testing.T) {
	original := &AiozAPIError{
		StatusCode: 403,
		ErrorCode:  "FORBIDDEN",
		Message:    "Access denied",
		Endpoint:   "/api-key/model/delete",
		Method:     "DELETE",
	}

	var err error = original

	var apiErr *AiozAPIError
	require.True(t, errors.As(err, &apiErr))
	assert.Equal(t, 403, apiErr.StatusCode)
	assert.Equal(t, "FORBIDDEN", apiErr.ErrorCode)
	assert.Equal(t, "Access denied", apiErr.Message)
}

func TestAiozAPIError_JSONMarshal(t *testing.T) {
	err := &AiozAPIError{
		StatusCode: 400,
		ErrorCode:  "VALIDATION_ERROR",
		Message:    "Invalid input",
		Endpoint:   "/api-key/model/create",
		Method:     "POST",
		RequestID:  "req-xyz",
	}

	data, jsonErr := json.Marshal(err)
	require.NoError(t, jsonErr)

	var parsed AiozAPIError
	require.NoError(t, json.Unmarshal(data, &parsed))

	assert.Equal(t, 400, parsed.StatusCode)
	assert.Equal(t, "VALIDATION_ERROR", parsed.ErrorCode)
	assert.Equal(t, "Invalid input", parsed.Message)
	assert.Equal(t, "req-xyz", parsed.RequestID)
}

func TestNewAPIError(t *testing.T) {
	err := NewAPIError(502, "BAD_GATEWAY", "Upstream error", "/api-key/model/list", "POST", "req-123")

	assert.Equal(t, 502, err.StatusCode)
	assert.Equal(t, "BAD_GATEWAY", err.ErrorCode)
	assert.Equal(t, "Upstream error", err.Message)
	assert.Equal(t, "/api-key/model/list", err.Endpoint)
	assert.Equal(t, "POST", err.Method)
	assert.Equal(t, "req-123", err.RequestID)
}
