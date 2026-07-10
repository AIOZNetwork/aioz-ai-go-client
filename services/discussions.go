package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/comment"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/discussion"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/reaction"
)

// DiscussionsService provides access to Service wrapper for all discussion-related API operations.-related API operations.
type DiscussionsService struct {
	Comment comment.ClientService
	Discussion discussion.ClientService
	Reaction reaction.ClientService
}

// NewDiscussionsService creates a DiscussionsService from the generated client.
func NewDiscussionsService(c *apiclient.AiozaiPlatform) *DiscussionsService {
	return &DiscussionsService{
		Comment: c.Comment,
		Discussion: c.Discussion,
		Reaction: c.Reaction,
	}
}
