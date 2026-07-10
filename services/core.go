package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/core"
)

// CoreService provides access to Service wrapper for all core-related API operations.-related API operations.
type CoreService struct {
	Core core.ClientService
}

// NewCoreService creates a CoreService from the generated client.
func NewCoreService(c *apiclient.AiozaiPlatform) *CoreService {
	return &CoreService{
		Core: c.Core,
	}
}
