package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/organization"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/organization_wallet"
)

// OrganizationsService provides access to organization-related API operations.
type OrganizationsService struct {
	Organization organization.ClientService
	Wallet       organization_wallet.ClientService
}

// NewOrganizationsService creates an OrganizationsService from the generated client.
func NewOrganizationsService(c *apiclient.AiozaiPlatform) *OrganizationsService {
	return &OrganizationsService{
		Organization: c.Organization,
		Wallet:       c.OrganizationWallet,
	}
}
