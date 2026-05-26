package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_organization"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/api_key_organization_wallet"
)

// OrganizationsService provides access to organization-related API operations.
type OrganizationsService struct {
	Organization api_key_organization.ClientService
	Wallet       api_key_organization_wallet.ClientService
}

// NewOrganizationsService creates an OrganizationsService from the generated client.
func NewOrganizationsService(c *apiclient.AiozaiPlatform) *OrganizationsService {
	return &OrganizationsService{
		Organization: c.APIKeyOrganization,
		Wallet:       c.APIKeyOrganizationWallet,
	}
}
