package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/task"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/dependency"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/offer"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/packages"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/platform_task"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/reaction"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/search"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/training_task"
)

// TaskService provides access to core platform API operations.
type TaskService struct {
	Task         task.ClientService
	Dependency   dependency.ClientService
	Offer        offer.ClientService
	Package      packages.ClientService
	PlatformTask platform_task.ClientService
	Reaction     reaction.ClientService
	Search       search.ClientService
	TrainingTask training_task.ClientService
}

// NewTaskService creates a TaskService from the generated client.
func NewTaskService(c *apiclient.AiozaiPlatform) *TaskService {
	return &TaskService{
		Task:         c.Task,
		Dependency:   c.Dependency,
		Offer:        c.Offer,
		Package:      c.Packages,
		PlatformTask: c.PlatformTask,
		Reaction:     c.Reaction,
		Search:       c.Search,
		TrainingTask: c.TrainingTask,
	}
}
