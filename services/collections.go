package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/collection"
)

// CollectionsService provides access to Service wrapper for all collection-related API operations.-related API operations.
type CollectionsService struct {
	Collection collection.ClientService
}

// NewCollectionsService creates a CollectionsService from the generated client.
func NewCollectionsService(c *apiclient.AiozaiPlatform) *CollectionsService {
	return &CollectionsService{
		Collection: c.Collection,
	}
}
