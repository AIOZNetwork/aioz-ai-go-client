package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_user"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_user_public_key"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_user_voucher"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_user_wallet"
)

// UsersService provides access to user-related API operations.
type UsersService struct {
	User      api_key_user.ClientService
	PublicKey api_key_user_public_key.ClientService
	Voucher   api_key_user_voucher.ClientService
	Wallet    api_key_user_wallet.ClientService
}

// NewUsersService creates a UsersService from the generated client.
func NewUsersService(c *apiclient.AiozaiPlatform) *UsersService {
	return &UsersService{
		User:      c.APIKeyUser,
		PublicKey: c.APIKeyUserPublicKey,
		Voucher:   c.APIKeyUserVoucher,
		Wallet:    c.APIKeyUserWallet,
	}
}
