package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_comments"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_discussion"
)

// DiscussionsService provides access to discussion-related API operations.
type DiscussionsService struct {
	Discussion api_key_discussion.ClientService
	Comments   api_key_comments.ClientService
}

// NewDiscussionsService creates a DiscussionsService from the generated client.
func NewDiscussionsService(c *apiclient.AiozaiPlatform) *DiscussionsService {
	return &DiscussionsService{
		Discussion: c.APIKeyDiscussion,
		Comments:   c.APIKeyComments,
	}
}
