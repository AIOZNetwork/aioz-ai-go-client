package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_dependency"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_offer"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_package"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_platform_task"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_reaction"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_search"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_training_task"
)

// CoreService provides access to core platform API operations.
type CoreService struct {
	APIKey       api_key.ClientService
	Dependency   api_key_dependency.ClientService
	Offer        api_key_offer.ClientService
	Package      api_key_package.ClientService
	PlatformTask api_key_platform_task.ClientService
	Reaction     api_key_reaction.ClientService
	Search       api_key_search.ClientService
	TrainingTask api_key_training_task.ClientService
}

// NewCoreService creates a CoreService from the generated client.
func NewCoreService(c *apiclient.AiozaiPlatform) *CoreService {
	return &CoreService{
		APIKey:       c.APIKey,
		Dependency:   c.APIKeyDependency,
		Offer:        c.APIKeyOffer,
		Package:      c.APIKeyPackage,
		PlatformTask: c.APIKeyPlatformTask,
		Reaction:     c.APIKeyReaction,
		Search:       c.APIKeySearch,
		TrainingTask: c.APIKeyTrainingTask,
	}
}
