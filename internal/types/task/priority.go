package task

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"strings"
)

type Priority int

const (
	PriorityNone   Priority = 0
	PriorityLow    Priority = 1
	PriorityMedium Priority = 3
	PriorityHigh   Priority = 5
)

var PriorityCompletion = []cobra.Completion{
	cobra.CompletionWithDesc("none", "No task"),
	cobra.CompletionWithDesc("low", "Low task"),
	cobra.CompletionWithDesc("medium", "Medium task"),
	cobra.CompletionWithDesc("high", "High task"),
}

var PriorityCompletionFunc = cobra.FixedCompletions(PriorityCompletion, cobra.ShellCompDirectiveNoFileComp)

var (
	NonePriorityColor   = color.HEX("#C6C6C6").C256()
	LowPriorityColor    = color.HEX("#4772F9").C256()
	MediumPriorityColor = color.HEX("#FAA80B").C256()
	HighPriorityColor   = color.HEX("#D52B24").C256()
)

var priorityMap = map[string]Priority{
	"none":   PriorityNone,
	"low":    PriorityLow,
	"medium": PriorityMedium,
	"high":   PriorityHigh,
}

func (p *Priority) UnmarshalJSON(data []byte) error {
	var priority int
	if err := json.Unmarshal(data, &priority); err != nil {
		return err
	}
	switch priority {
	case int(PriorityNone), int(PriorityLow), int(PriorityMedium), int(PriorityHigh):
		*p = Priority(priority)
	default:
		*p = PriorityNone
	}
	return nil
}

func (p Priority) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(p))
}

func (p Priority) String() string {
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

func (p *Priority) Set(value string) error {
	priority, ok := priorityMap[strings.ToLower(value)]
	if !ok {
		return fmt.Errorf("invalid task: %s", value)
	}

	*p = priority
	return nil
}

func (p *Priority) Type() string {
	return "Priority"
}
