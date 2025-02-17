package types

import (
	"encoding/json"
	"github.com/gookit/color"
)

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

func (s TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(s))
}

func (s TaskStatus) String() string {
	switch s {
	case StatusComplete:
		return color.Green.Sprint("☑")
	case StatusNormal:
		return color.White.Sprint("☐")
	default:
		return color.Red.Sprint("☒")
	}
}
