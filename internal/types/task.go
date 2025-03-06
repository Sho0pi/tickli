package types

import "github.com/sho0pi/tickli/internal/types/task"

type Task struct {
	ID            string          `json:"id"`
	ProjectID     string          `json:"projectId"`
	Title         string          `json:"title"`
	IsAllDay      bool            `json:"isAllDay"`
	CompletedTime TickTickTime    `json:"completedTime,omitzero"`
	Content       string          `json:"content"`
	Desc          string          `json:"desc"`
	DueDate       TickTickTime    `json:"dueDate,omitzero"`
	Items         []ChecklistItem `json:"items"`
	Priority      task.Priority   `json:"priority"`
	Reminders     []string        `json:"reminders"`
	RepeatFlag    string          `json:"repeatFlag"`
	SortOrder     int64           `json:"sortOrder"`
	StartDate     TickTickTime    `json:"startDate,omitzero"`
	Status        task.Status     `json:"status"`
	TimeZone      string          `json:"timeZone"`
	Tags          []string        `json:"tags"`
}

type ChecklistItem struct {
	ID            string       `json:"id"`
	Title         string       `json:"title"`
	Status        int          `json:"status"`
	CompletedTime int64        `json:"completedTime"`
	IsAllDay      bool         `json:"isAllDay"`
	SortOrder     int64        `json:"sortOrder"`
	StartDate     TickTickTime `json:"startDate"`
	TimeZone      string       `json:"timeZone"`
}
