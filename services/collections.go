package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_collection"
)

// CollectionsService provides access to collection-related API operations.
type CollectionsService struct {
	Collection api_key_collection.ClientService
}

// NewCollectionsService creates a CollectionsService from the generated client.
func NewCollectionsService(c *apiclient.AiozaiPlatform) *CollectionsService {
	return &CollectionsService{
		Collection: c.APIKeyCollection,
	}
}
