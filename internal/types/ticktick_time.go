package types

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"time"
)

type TickTickTime time.Time

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

func (t TickTickTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format("2006-01-02T15:04:05-0700"))
}

func (t TickTickTime) String() string {
	return time.Time(t).Format("Monday 2006-01-02 15:04:05")
}

func (t TickTickTime) Humanize() string {
	return humanize.Time(time.Time(t))
}
