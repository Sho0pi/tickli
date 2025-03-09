package utils

import (
	"github.com/sho0pi/naturaltime"
	"time"
)

var DefaultDuration = 1 * time.Hour

func ParseTimeExpression(expr string) (*naturaltime.Range, error) {
	p, err := naturaltime.New()
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()
	r, err := p.ParseRange(expr, currentTime)
	if err != nil {
		return nil, err
	}

	if !r.IsAllDay() {
		return r, nil
	}
	// Checks if the user just specified the date but no the time.
	if r.Start().Hour() == currentTime.Hour() && r.Start().Minute() == currentTime.Minute() && r.Start().Second() == currentTime.Second() {
		return r, nil
	}
	r.Duration = DefaultDuration
	return r, nil
}
