package types

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"strings"
)

type TaskPriority int

const (
	PriorityNone   TaskPriority = 0
	PriorityLow    TaskPriority = 1
	PriorityMedium TaskPriority = 3
	PriorityHigh   TaskPriority = 5
)

func RegisterTaskPriorityCompletions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return priorityMapKeys(), cobra.ShellCompDirectiveDefault
}

func priorityMapKeys() []string {
	keys := make([]string, 0, len(priorityMap))
	for k := range priorityMap {
		keys = append(keys, k)
	}
	return keys
}

var (
	NonePriorityColor   = color.HEX("#C6C6C6").C256()
	LowPriorityColor    = color.HEX("#4772F9").C256()
	MediumPriorityColor = color.HEX("#FAA80B").C256()
	HighPriorityColor   = color.HEX("#D52B24").C256()
)

var priorityMap = map[string]TaskPriority{
	"none":   PriorityNone,
	"low":    PriorityLow,
	"medium": PriorityMedium,
	"high":   PriorityHigh,
}

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

func (p *TaskPriority) Set(value string) error {
	priority, ok := priorityMap[strings.ToLower(value)]
	if !ok {
		return fmt.Errorf("invalid priority: %s", value)
	}

	*p = priority
	return nil
}

func (p *TaskPriority) Type() string {
	return "TaskPriority"
}
