package types

import (
	"encoding/json"
	"github.com/spf13/cobra"
)

// ViewMode describes the view mode of the project, by default will be "list"
type ViewMode string

const (
	ViewModeList     ViewMode = "list"
	ViewModeKanban   ViewMode = "kanban"
	ViewModeTimeline ViewMode = "timeline"
)

func RegisterViewModeCompletions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{string(ViewModeList), string(ViewModeTimeline), string(ViewModeKanban)}, cobra.ShellCompDirectiveDefault
}

func (vm *ViewMode) UnmarshalJSON(data []byte) error {
	var viewMode string
	if err := json.Unmarshal(data, &viewMode); err != nil {
		return err
	}
	switch viewMode {
	case string(ViewModeList), string(ViewModeKanban), string(ViewModeTimeline):
		*vm = ViewMode(viewMode)
	default:
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
	switch s {
	case string(ViewModeList), string(ViewModeKanban), string(ViewModeTimeline):
		*vm = ViewMode(s)
	default:
		*vm = ViewModeList
	}
	return nil
}

func (vm *ViewMode) Type() string {
	return "ViewMode"
}
