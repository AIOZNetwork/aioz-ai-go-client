package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/competition"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/competition_leaderboard"
)

// CompetitionsService provides access to competition-related API operations.
type CompetitionsService struct {
	Competition competition.ClientService
	Leaderboard competition_leaderboard.ClientService
}

// NewCompetitionsService creates a CompetitionsService from the generated client.
func NewCompetitionsService(c *apiclient.AiozaiPlatform) *CompetitionsService {
	return &CompetitionsService{
		Competition: c.Competition,
		Leaderboard: c.CompetitionLeaderboard,
	}
}
