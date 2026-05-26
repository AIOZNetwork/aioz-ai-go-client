package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_model"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_model_api_key"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_model_package"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_model_playground"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_model_reviews"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_model_setting"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_model_training"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_model_verify"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_model_versioning"
)

// ModelsService provides access to model-related API operations.
type ModelsService struct {
	Model        api_key_model.ClientService
	ModelAPIKey  api_key_model_api_key.ClientService
	ModelPackage api_key_model_package.ClientService
	Playground   api_key_model_playground.ClientService
	Reviews      api_key_model_reviews.ClientService
	Setting      api_key_model_setting.ClientService
	Training     api_key_model_training.ClientService
	Verify       api_key_model_verify.ClientService
	Versioning   api_key_model_versioning.ClientService
}

// NewModelsService creates a ModelsService from the generated client.
func NewModelsService(c *apiclient.AiozaiPlatform) *ModelsService {
	return &ModelsService{
		Model:        c.APIKeyModel,
		ModelAPIKey:  c.APIKeyModelAPIKey,
		ModelPackage: c.APIKeyModelPackage,
		Playground:   c.APIKeyModelPlayground,
		Reviews:      c.APIKeyModelReviews,
		Setting:      c.APIKeyModelSetting,
		Training:     c.APIKeyModelTraining,
		Verify:       c.APIKeyModelVerify,
		Versioning:   c.APIKeyModelVersioning,
	}
}
