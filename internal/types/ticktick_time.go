package types

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/ijt/go-anytime"
	"time"
)

type TickTickTime time.Time

func (t *TickTickTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
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

func GetRangeFromString(timeRange string) (start TickTickTime, end TickTickTime, err error) {
	r, err := anytime.ParseRange(timeRange, time.Now())
	if err != nil {
		return TickTickTime{}, TickTickTime{}, err
	}
	return TickTickTime(r.Start()), TickTickTime(r.End()), nil
}
