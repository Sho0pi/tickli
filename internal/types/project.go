package types

// InboxProject the Inbox project representation (cause is not returned by the api)
var InboxProject = Project{
	ID:        "inbox",
	Name:      "inbox",
	Color:     DefaultColor,
	SortOrder: 0,
	Closed:    false,
	Kind:      KindInbox,
	ViewMode:  ViewModeList,
}

type Project struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Color      ProjectColor `json:"color"`
	SortOrder  int64        `json:"sortOrder"`
	Closed     bool         `json:"closed"`
	GroupID    string       `json:"groupId"`
	ViewMode   ViewMode     `json:"viewMode"`
	Permission string       `json:"permission"`
	Kind       ProjectKind  `json:"kind"`
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
