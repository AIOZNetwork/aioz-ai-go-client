package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/user"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/user_payment"
)

// UsersService provides access to Service wrapper for all user-related API operations.-related API operations.
type UsersService struct {
	User user.ClientService
	Payment user_payment.ClientService
}

// NewUsersService creates a UsersService from the generated client.
func NewUsersService(c *apiclient.AiozaiPlatform) *UsersService {
	return &UsersService{
		User: c.User,
		Payment: c.UserPayment,
	}
}
