package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_dataset"
)

// DatasetsService provides access to dataset-related API operations.
type DatasetsService struct {
	Dataset api_key_dataset.ClientService
}

// NewDatasetsService creates a DatasetsService from the generated client.
func NewDatasetsService(c *apiclient.AiozaiPlatform) *DatasetsService {
	return &DatasetsService{
		Dataset: c.APIKeyDataset,
	}
}
