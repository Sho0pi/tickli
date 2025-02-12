package api

import "time"

// TODO: fix time problem
type Project struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	SortOrder  int64  `json:"sortOrder"`
	Closed     bool   `json:"closed"`
	GroupID    string `json:"groupId"`
	ViewMode   string `json:"viewMode"`
	Permission string `json:"permission"`
	Kind       string `json:"kind"`
}

type ChecklistItem struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Status        int       `json:"status"`
	CompletedTime time.Time `json:"completedTime"`
	IsAllDay      bool      `json:"isAllDay"`
	SortOrder     int64     `json:"sortOrder"`
	StartDate     time.Time `json:"startDate"`
	TimeZone      string    `json:"timeZone"`
}

type Task struct {
	ID            string          `json:"id"`
	ProjectID     string          `json:"projectId"`
	Title         string          `json:"title"`
	IsAllDay      bool            `json:"isAllDay"`
	CompletedTime *time.Time      `json:"completedTime"`
	Content       string          `json:"content"`
	Desc          string          `json:"desc"`
	DueDate       *time.Time      `json:"dueDate"`
	Items         []ChecklistItem `json:"items"`
	Priority      int             `json:"priority"`
	Reminders     []string        `json:"reminders"`
	RepeatFlag    string          `json:"repeatFlag"`
	SortOrder     int64           `json:"sortOrder"`
	StartDate     *time.Time      `json:"startDate"`
	Status        int             `json:"status"`
	TimeZone      string          `json:"timeZone"`
}

type Column struct {
	ID        string `json:"id"`
	ProjectID string `json:"projectId"`
	Name      string `json:"name"`
	SortOrder int64  `json:"sortOrder"`
}

type ProjectData struct {
	Project Project  `json:"project"`
	Tasks   []Task   `json:"tasks"`
	Columns []Column `json:"columns"`
}
