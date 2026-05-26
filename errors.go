package aiozai

import "fmt"

// AiozAPIError represents a structured error from the AIOZ AI API.
type AiozAPIError struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
	Endpoint   string `json:"endpoint"`
	Method     string `json:"method"`
	RequestID  string `json:"requestId,omitempty"`
}

// Error implements the error interface.
func (e *AiozAPIError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("[%d] %s: %s (%s %s, request-id: %s)",
			e.StatusCode, e.ErrorCode, e.Message, e.Method, e.Endpoint, e.RequestID)
	}
	return fmt.Sprintf("[%d] %s: %s (%s %s)",
		e.StatusCode, e.ErrorCode, e.Message, e.Method, e.Endpoint)
}

// NewAPIError creates a new AiozAPIError.
func NewAPIError(statusCode int, errorCode, message, endpoint, method, requestID string) *AiozAPIError {
	return &AiozAPIError{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Message:    message,
		Endpoint:   endpoint,
		Method:     method,
		RequestID:  requestID,
	}
}
