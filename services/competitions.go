package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/competition"
)

// CompetitionsService provides access to Service wrapper for all competition-related API operations.-related API operations.
type CompetitionsService struct {
	Competition competition.ClientService
}

// NewCompetitionsService creates a CompetitionsService from the generated client.
func NewCompetitionsService(c *apiclient.AiozaiPlatform) *CompetitionsService {
	return &CompetitionsService{
		Competition: c.Competition,
	}
}
