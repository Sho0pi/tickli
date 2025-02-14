package api

import (
	"fmt"
	"github.com/sho0pi/tickli/internal/config"
)

func GetProjects() ([]Project, error) {
	token, err := config.LoadToken()
	if err != nil {
		return nil, fmt.Errorf("falied to load token: %w", err)
	}

	client := NewClient(token)
	projects, err := client.ListProjects()
	if err != nil {
		return nil, fmt.Errorf("falied to list projects: %w", err)
	}

	return projects, nil
}

func GetTasks(projectID string) ([]Task, error) {
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
