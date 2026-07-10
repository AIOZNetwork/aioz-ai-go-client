package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model"
)

// ModelsService provides access to Service wrapper for all model-related API operations.-related API operations.
type ModelsService struct {
	Model model.ClientService
}

// NewModelsService creates a ModelsService from the generated client.
func NewModelsService(c *apiclient.AiozaiPlatform) *ModelsService {
	return &ModelsService{
		Model: c.Model,
	}
}
