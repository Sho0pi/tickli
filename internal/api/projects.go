package api

import (
	"fmt"
	"github.com/sho0pi/tickli/internal/config"
	"github.com/sho0pi/tickli/internal/types"
)

func GetProjects() ([]types.Project, error) {
	token, err := config.LoadToken()
	if err != nil {
		return nil, fmt.Errorf("falied to load token: %w", err)
	}

	client := NewClient(token)
	projects, err := client.ListProjects()
	if err != nil {
		return nil, fmt.Errorf("falied to list projects: %w", err)
	}

	// Adds the default InboxProject - not appears by default
	projects = append(projects, types.InboxProject)

	return projects, nil
}

func GetTasks(projectID string) ([]types.Task, error) {
	token, err := config.LoadToken()
	if err != nil {
		return nil, fmt.Errorf("falied to load token: %w", err)
	}

	client := NewClient(token)
	tasks, err := client.ListTasks(projectID)
	if err != nil {
		return nil, fmt.Errorf("falied to list tasks: %w", err)
	}

	return tasks, nil
}
