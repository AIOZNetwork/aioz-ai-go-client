package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/public"
)

// PublicService provides access to Service for public (unauthenticated) API operations.-related API operations.
type PublicService struct {
	Public public.ClientService
}

// NewPublicService creates a PublicService from the generated client.
func NewPublicService(c *apiclient.AiozaiPlatform) *PublicService {
	return &PublicService{
		Public: c.Public,
	}
}
