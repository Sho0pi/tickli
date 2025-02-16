package api

import (
	"fmt"
	"github.com/sho0pi/tickli/internal/config"
)

// InboxProject the Inbox project representation (cause is not returned by the api)
var InboxProject = Project{
	ID:        "inbox",
	Name:      "inbox",
	Color:     DefaultColor,
	SortOrder: 0,
	Closed:    false,
	Kind:      KindInbox,
}

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

	// Adds the default InboxProject - not appears by default
	projects = append(projects, InboxProject)

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
