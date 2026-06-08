package aiozai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/services"
)

// Client is the AIOZ AI SDK client.
type Client struct {
	config    *ClientConfig
	raw       *apiclient.AiozaiPlatform
	transport *httptransport.Runtime
}

// NewClient creates a new SDK client with the given options.
func NewClient(opts ...Option) (*Client, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	parsed, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL %q: %w", cfg.BaseURL, err)
	}

	host := parsed.Host
	basePath := parsed.Path
	schemes := []string{parsed.Scheme}
	if parsed.Scheme == "" {
		schemes = []string{"https"}
	}

	// Create base HTTP transport
	var baseRT http.RoundTripper
	if cfg.HTTPClient != nil {
		baseRT = cfg.HTTPClient.Transport
	}
	if baseRT == nil {
		baseRT = http.DefaultTransport
	}

	// Wrap with retry transport
	rt := newRetryTransport(baseRT, cfg.RetryConfig)

	// Wrap with error-converting transport
	rt = &errorTransport{next: rt}

	// Create HTTP client with timeout
	httpClient := &http.Client{
		Transport: rt,
		Timeout:   cfg.Timeout,
	}

	transport := httptransport.NewWithClient(host, basePath, schemes, httpClient)

	// Set API Key auth if API key provided
	if cfg.APIKey != "" {
		transport.DefaultAuthentication = httptransport.APIKeyAuth("x-api-key", "header", cfg.APIKey)
	}

	raw := apiclient.New(transport, strfmt.Default)

	return &Client{
		config:    cfg,
		raw:       raw,
		transport: transport,
	}, nil
}

// Raw returns the underlying go-swagger generated client for direct access.
func (c *Client) Raw() *apiclient.AiozaiPlatform {
	return c.raw
}

// Config returns the resolved client configuration.
func (c *Client) Config() *ClientConfig {
	return c.config
}

// Models returns the models service for AI model management operations.
func (c *Client) Models() *services.ModelsService {
	return services.NewModelsService(c.raw)
}

// Datasets returns the datasets service for dataset management operations.
func (c *Client) Datasets() *services.DatasetsService {
	return services.NewDatasetsService(c.raw)
}

// Competitions returns the competitions service for competition operations.
func (c *Client) Competitions() *services.CompetitionsService {
	return services.NewCompetitionsService(c.raw)
}

// Collections returns the collections service for curated collection operations.
func (c *Client) Collections() *services.CollectionsService {
	return services.NewCollectionsService(c.raw)
}

// Discussions returns the discussions service for discussion and comment operations.
func (c *Client) Discussions() *services.DiscussionsService {
	return services.NewDiscussionsService(c.raw)
}

// Notifications returns the notifications service.
func (c *Client) Notifications() *services.NotificationsService {
	return services.NewNotificationsService(c.raw)
}

// Organizations returns the organizations service for org management operations.
func (c *Client) Organizations() *services.OrganizationsService {
	return services.NewOrganizationsService(c.raw)
}

// Repositories returns the repositories service for repository operations.
func (c *Client) Repositories() *services.RepositoriesService {
	return services.NewRepositoriesService(c.raw)
}

// Storage returns the storage service for upload and storage operations.
func (c *Client) Storage() *services.StorageService {
	return services.NewStorageService(c.raw)
}

// Users returns the users service for user management operations.
func (c *Client) Users() *services.UsersService {
	return services.NewUsersService(c.raw)
}

// Public returns the public service for unauthenticated endpoints.
func (c *Client) Public() *services.PublicService {
	return services.NewPublicService(c.raw)
}

// Core returns the core service for root endpoints and minor services.
func (c *Client) Core() *services.TaskService {
	return services.NewTaskService(c.raw)
}

// errorTransport wraps an http.RoundTripper and converts non-2xx responses
// to AiozAPIError.
type errorTransport struct {
	next http.RoundTripper
}

func (t *errorTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.next.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return resp, nil
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	apiErr := &AiozAPIError{
		StatusCode: resp.StatusCode,
		Method:     req.Method,
		Endpoint:   req.URL.Path,
		RequestID:  resp.Header.Get("X-Request-Id"),
	}

	// Try to parse error response body
	var errBody struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}
	if json.Unmarshal(body, &errBody) == nil {
		apiErr.ErrorCode = errBody.Code
		apiErr.Message = errBody.Message
		if apiErr.Message == "" {
			apiErr.Message = errBody.Error
		}
	}

	if apiErr.ErrorCode == "" {
		apiErr.ErrorCode = http.StatusText(resp.StatusCode)
	}
	if apiErr.Message == "" {
		apiErr.Message = strings.TrimSpace(string(body))
	}

	return nil, apiErr
}
