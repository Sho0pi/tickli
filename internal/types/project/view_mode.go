package project

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

// ViewMode represents the different view modes available in the project, by default will be "list"
type ViewMode string

const (
	ViewModeList     ViewMode = "list"
	ViewModeKanban   ViewMode = "kanban"
	ViewModeTimeline ViewMode = "timeline"
)

var ViewModeCompletion = []cobra.Completion{
	cobra.CompletionWithDesc(ViewModeList.String(), "List view mode"),
	cobra.CompletionWithDesc(ViewModeKanban.String(), "Kanban view mode"),
	cobra.CompletionWithDesc(ViewModeTimeline.String(), "Timeline view mode"),
}

var ViewModeCompletionFunc = cobra.FixedCompletions(ViewModeCompletion, cobra.ShellCompDirectiveNoFileComp)

func (vm *ViewMode) UnmarshalJSON(data []byte) error {
	var viewMode string
	if err := json.Unmarshal(data, &viewMode); err != nil {
		return err
	}
	if isValidViewMode(viewMode) {
		*vm = ViewMode(viewMode)
	} else {
		*vm = ViewModeList
	}
	return nil
}

func (vm ViewMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(vm))
}

func (vm ViewMode) String() string {
	return string(vm)
}

func (vm *ViewMode) Set(s string) error {
	if isValidViewMode(s) {
		*vm = ViewMode(s)
		return nil
	}
	return fmt.Errorf("invalid view mode %q", s)
}

func isValidViewMode(mode string) bool {
	switch mode {
	case string(ViewModeList), string(ViewModeKanban), string(ViewModeTimeline):
		return true
	default:
		return false
	}
}

func (vm *ViewMode) Type() string {
	return "ViewMode"
}
