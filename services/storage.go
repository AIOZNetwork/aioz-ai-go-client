package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_storage"
)

// StorageService provides access to storage-related API operations.
type StorageService struct {
	Storage api_key_storage.ClientService
}

// NewStorageService creates a StorageService from the generated client.
func NewStorageService(c *apiclient.AiozaiPlatform) *StorageService {
	return &StorageService{
		Storage: c.APIKeyStorage,
	}
}
