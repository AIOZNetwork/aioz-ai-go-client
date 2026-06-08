package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/comments"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/discussion"
)

// DiscussionsService provides access to discussion-related API operations.
type DiscussionsService struct {
	Discussion discussion.ClientService
	Comments   comments.ClientService
}

// NewDiscussionsService creates a DiscussionsService from the generated client.
func NewDiscussionsService(c *apiclient.AiozaiPlatform) *DiscussionsService {
	return &DiscussionsService{
		Discussion: c.Discussion,
		Comments:   c.Comments,
	}
}
