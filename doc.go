// Package aiozai provides a Go client SDK for the AIOZ AI API.
//
// The SDK is auto-generated from the OpenAPI specification and provides
// typed access to all endpoints with all model definitions.
//
// # Quick Start
//
//	client, err := aiozai.NewClient(
//	    aiozai.WithAPIKey(os.Getenv("AIOZ_AI_API_KEY")),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// # Features
//
//   - Full type safety for all API endpoints and models
//   - API Key configured once at client creation
//   - Automatic retry with exponential backoff for transient failures
//   - Configurable timeouts for standard and upload operations
//   - Domain-grouped service access (Models, Datasets, Competitions, etc.)
//
// # Authentication
//
// The SDK uses API Key. Set your API key when creating
// the client with WithAPIKey(). All api-key/* endpoints automatically
// include the x-api-key header. Public endpoints (public/*) do not
// require authentication.
//
// # Error Handling
//
// API errors are returned as *AiozAPIError which implements the error
// interface. Use errors.As() to check for API-specific errors:
//
//	var apiErr *aiozai.AiozAPIError
//	if errors.As(err, &apiErr) {
//	    log.Printf("API error %d: %s", apiErr.StatusCode, apiErr.Message)
//	}
package aiozai
