package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/public"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/public_collection"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/public_competition"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/public_dataset"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/public_discussion"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/public_medals"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/public_model"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/public_organization"
)

// PublicService provides access to public (unauthenticated) API operations.
type PublicService struct {
	Public       public.ClientService
	Collection   public_collection.ClientService
	Competition  public_competition.ClientService
	Dataset      public_dataset.ClientService
	Discussion   public_discussion.ClientService
	Medals       public_medals.ClientService
	Model        public_model.ClientService
	Organization public_organization.ClientService
}

// NewPublicService creates a PublicService from the generated client.
func NewPublicService(c *apiclient.AiozaiPlatform) *PublicService {
	return &PublicService{
		Public:       c.Public,
		Collection:   c.PublicCollection,
		Competition:  c.PublicCompetition,
		Dataset:      c.PublicDataset,
		Discussion:   c.PublicDiscussion,
		Medals:       c.PublicMedals,
		Model:        c.PublicModel,
		Organization: c.PublicOrganization,
	}
}
