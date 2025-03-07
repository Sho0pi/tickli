package types

import "github.com/sho0pi/tickli/internal/types/project"

// InboxProject the Inbox project representation (cause is not returned by the api)
var InboxProject = Project{
	ID:        "inbox",
	Name:      "📥Inbox",
	Color:     project.DefaultColor,
	SortOrder: 0,
	Closed:    false,
	Kind:      project.KindTask,
	ViewMode:  project.ViewModeList,
}

var NullProject = Project{}

type Project struct {
	ID         string           `json:"id"`
	Name       string           `json:"name"`
	Color      project.Color    `json:"color"`
	SortOrder  int64            `json:"sortOrder"`
	Closed     bool             `json:"closed"`
	GroupID    string           `json:"groupId"`
	ViewMode   project.ViewMode `json:"viewMode"`
	Permission string           `json:"permission"`
	Kind       project.Kind     `json:"kind"`
}

type ProjectData struct {
	Project Project  `json:"project"`
	Tasks   []Task   `json:"tasks"`
	Columns []Column `json:"columns"`
}

type Column struct {
	ID        string `json:"id"`
	ProjectID string `json:"projectId"`
	Name      string `json:"name"`
	SortOrder int64  `json:"sortOrder"`
}
