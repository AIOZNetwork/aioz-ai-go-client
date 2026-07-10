package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/organization"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/organization_payment"
)

// OrganizationsService provides access to Service for managing organizations and their payments.-related API operations.
type OrganizationsService struct {
	Organization organization.ClientService
	Payment organization_payment.ClientService
}

// NewOrganizationsService creates a OrganizationsService from the generated client.
func NewOrganizationsService(c *apiclient.AiozaiPlatform) *OrganizationsService {
	return &OrganizationsService{
		Organization: c.Organization,
		Payment: c.OrganizationPayment,
	}
}
