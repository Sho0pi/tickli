package types

import "encoding/json"

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

func (vm ViewMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(vm))
}

func (vm ViewMode) String() string {
	return string(vm)
}
