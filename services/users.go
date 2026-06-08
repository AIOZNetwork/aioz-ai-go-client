package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/user"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/user_public_key"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/user_voucher"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/user_wallet"
)

// UsersService provides access to user-related API operations.
type UsersService struct {
	User      user.ClientService
	PublicKey user_public_key.ClientService
	Voucher   user_voucher.ClientService
	Wallet    user_wallet.ClientService
}

// NewUsersService creates a UsersService from the generated client.
func NewUsersService(c *apiclient.AiozaiPlatform) *UsersService {
	return &UsersService{
		User:      c.User,
		PublicKey: c.UserPublicKey,
		Voucher:   c.UserVoucher,
		Wallet:    c.UserWallet,
	}
}
