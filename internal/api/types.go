package api

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"time"
)

type ProjectColor color.Color256

type ProjectKind string

const (
	KindTask    ProjectKind = "TASK"
	KindNote    ProjectKind = "NOTE"
	KindInbox   ProjectKind = "INBOX"
	KindUnknown ProjectKind = "UNKNOWN"
)

// ViewMode describes the view mode of the project, by default will be "list"
type ViewMode string

const (
	ViewModeList     = "list"
	ViewModeKanban   = "kanban"
	ViewModeTimeline = "timeline"
)

func (vm *ViewMode) UnmarshalJSON(data []byte) error {
	var viewMode string
	if err := json.Unmarshal(data, &viewMode); err != nil {
		return err
	}
	switch viewMode {
	case ViewModeList, ViewModeKanban, ViewModeTimeline:
		*vm = ViewMode(viewMode)
	default:
		*vm = ViewModeList
	}
	return nil
}

func (k *ProjectKind) UnmarshalJSON(data []byte) error {
	var kind string
	if err := json.Unmarshal(data, &kind); err != nil {
		return err
	}
	switch kind {
	case string(KindTask), string(KindNote), string(KindInbox), string(KindUnknown):
		*k = ProjectKind(kind)
	default:
		*k = KindUnknown
	}
	return nil
}

func (c *ProjectColor) UnmarshalJSON(data []byte) error {

	var colorStr string
	if err := json.Unmarshal(data, &colorStr); err != nil {
		return err
	}

	if colorStr == "" {
		*c = DefaultColor
	} else {
		*c = ProjectColor(color.HEX(colorStr).C256())
	}
	return nil
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

type TickTickTime time.Time

type TaskPriority int

const (
	PriorityNone   TaskPriority = 0
	PriorityLow    TaskPriority = 1
	PriorityMedium TaskPriority = 3
	PriorityHigh   TaskPriority = 5
)

func (p *TaskPriority) UnmarshalJSON(data []byte) error {
	var priority int
	if err := json.Unmarshal(data, &priority); err != nil {
		return err
	}
	switch priority {
	case int(PriorityNone), int(PriorityLow), int(PriorityMedium), int(PriorityHigh):
		*p = TaskPriority(priority)
	default:
		*p = PriorityNone
	}
	return nil
}

func (t *TickTickTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		fmt.Println(string(data))
		return err
	}

	ts, err := time.Parse("2006-01-02T15:04:05-0700", timeStr)
	if err != nil {
		return fmt.Errorf("invalid time format: %s", timeStr)
	}

	*t = TickTickTime(ts)
	return nil
}

type TaskStatus int

var (
	StatusNormal   TaskStatus = 0
	StatusComplete TaskStatus = 2
)

func (s *TaskStatus) UnmarshalJSON(data []byte) error {
	var status int
	if err := json.Unmarshal(data, &status); err != nil {
		return err
	}
	switch status {
	case int(StatusNormal), int(StatusComplete):
		*s = TaskStatus(status)
	default:
		*s = StatusNormal
	}
	return nil
}

type Task struct {
	ID            string          `json:"id"`
	ProjectID     string          `json:"projectId"`
	Title         string          `json:"title"`
	IsAllDay      bool            `json:"isAllDay"`
	CompletedTime TickTickTime    `json:"completedTime"`
	Content       string          `json:"content"`
	Desc          string          `json:"desc"`
	DueDate       TickTickTime    `json:"dueDate"`
	Items         []ChecklistItem `json:"items"`
	Priority      TaskPriority    `json:"priority"`
	Reminders     []string        `json:"reminders"`
	RepeatFlag    string          `json:"repeatFlag"`
	SortOrder     int64           `json:"sortOrder"`
	StartDate     TickTickTime    `json:"startDate"`
	Status        TaskStatus      `json:"status"`
	TimeZone      string          `json:"timeZone"`
	Tags          []string        `json:"tags"`
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

var DefaultColor = ProjectColor(color.HEX("#3694FE").C256())

func (t *TickTickTime) String() string {
	return time.Time(*t).Format("Monday 2006-01-02 15:04:05")
}

func (k *ProjectKind) GetEmoji() string {
	switch *k {
	case KindTask:
		return "📝Task"
	case KindNote:
		return "📖Note"
	case KindInbox:
		return "📥Inbox"
	default:
		return "🔧Unknown"
	}
}

var (
	NonePriorityColor   = color.HEX("#C6C6C6").C256()
	LowPriorityColor    = color.HEX("#4772F9").C256()
	MediumPriorityColor = color.HEX("#FAA80B").C256()
	HighPriorityColor   = color.HEX("#D52B24").C256()
)

func (p *TaskPriority) String() string {
	flag := "⚑"
	switch *p {
	case PriorityNone:
		flag = NonePriorityColor.Sprint(flag)
	case PriorityLow:
		flag = LowPriorityColor.Sprint(flag)
	case PriorityMedium:
		flag = MediumPriorityColor.Sprint(flag)
	case PriorityHigh:
		flag = HighPriorityColor.Sprint(flag)
	}

	return flag
}

func (p *Project) ColorSprint(a ...any) string {
	return color.Color256(p.Color).Sprint(a...)
}

func (s *TaskStatus) String() string {
	switch *s {
	case StatusComplete:
		return color.Green.Sprint("☑")
	case StatusNormal:
		return color.White.Sprint("☐")
	default:
		return color.Red.Sprint("☒")
	}
}
