package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_competition"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_competition_leaderboard"
)

// CompetitionsService provides access to competition-related API operations.
type CompetitionsService struct {
	Competition api_key_competition.ClientService
	Leaderboard api_key_competition_leaderboard.ClientService
}

// NewCompetitionsService creates a CompetitionsService from the generated client.
func NewCompetitionsService(c *apiclient.AiozaiPlatform) *CompetitionsService {
	return &CompetitionsService{
		Competition: c.APIKeyCompetition,
		Leaderboard: c.APIKeyCompetitionLeaderboard,
	}
}
