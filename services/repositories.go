package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/repository"
)

// RepositoriesService provides access to repository-related API operations.
type RepositoriesService struct {
	Repository repository.ClientService
}

// NewRepositoriesService creates a RepositoriesService from the generated client.
func NewRepositoriesService(c *apiclient.AiozaiPlatform) *RepositoriesService {
	return &RepositoriesService{
		Repository: c.Repository,
	}
}
