package types

import (
	"encoding/json"
	"github.com/gookit/color"
)

type TaskPriority int

const (
	PriorityNone   TaskPriority = 0
	PriorityLow    TaskPriority = 1
	PriorityMedium TaskPriority = 3
	PriorityHigh   TaskPriority = 5
)

var (
	NonePriorityColor   = color.HEX("#C6C6C6").C256()
	LowPriorityColor    = color.HEX("#4772F9").C256()
	MediumPriorityColor = color.HEX("#FAA80B").C256()
	HighPriorityColor   = color.HEX("#D52B24").C256()
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

func (p TaskPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(p))
}

func (p TaskPriority) String() string {
	flag := "⚑"
	switch p {
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
