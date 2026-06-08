package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model_api_key"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model_package"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model_playground"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model_reviews"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model_setting"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model_training"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model_verify"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/model_versioning"
)

// ModelsService provides access to model-related API operations.
type ModelsService struct {
	Model        model.ClientService
	ModelAPIKey  model_api_key.ClientService
	ModelPackage model_package.ClientService
	Playground   model_playground.ClientService
	Reviews      model_reviews.ClientService
	Setting      model_setting.ClientService
	Training     model_training.ClientService
	Verify       model_verify.ClientService
	Versioning   model_versioning.ClientService
}

// NewModelsService creates a ModelsService from the generated client.
func NewModelsService(c *apiclient.AiozaiPlatform) *ModelsService {
	return &ModelsService{
		Model:        c.Model,
		ModelAPIKey:  c.ModelAPIKey,
		ModelPackage: c.ModelPackage,
		Playground:   c.ModelPlayground,
		Reviews:      c.ModelReviews,
		Setting:      c.ModelSetting,
		Training:     c.ModelTraining,
		Verify:       c.ModelVerify,
		Versioning:   c.ModelVersioning,
	}
}
