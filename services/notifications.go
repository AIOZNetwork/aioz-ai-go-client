package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/notification"
)

// NotificationsService provides access to notification-related API operations.
type NotificationsService struct {
	Notification notification.ClientService
}

// NewNotificationsService creates a NotificationsService from the generated client.
func NewNotificationsService(c *apiclient.AiozaiPlatform) *NotificationsService {
	return &NotificationsService{
		Notification: c.Notification,
	}
}
