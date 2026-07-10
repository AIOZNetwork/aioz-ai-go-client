package services

import (
	apiclient "github.com/AIOZNetwork/aioz-ai-go-client/generated/client"
	"github.com/AIOZNetwork/aioz-ai-go-client/generated/client/task"
)

// TasksService provides access to Service wrapper for all task-related API operations.-related API operations.
type TasksService struct {
	Task task.ClientService
}

// NewTasksService creates a TasksService from the generated client.
func NewTasksService(c *apiclient.AiozaiPlatform) *TasksService {
	return &TasksService{
		Task: c.Task,
	}
}
