package goWallabag

import (
	"time"
)

const TimeFormat = "2006-01-02T15:04:05-0700"

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(timeByte []byte) error {
	timeStr := string(timeByte)
	timeStr = timeStr[1 : len(timeStr)-1]

	newTime, err := time.Parse(TimeFormat, timeStr)
	if err != nil {
		return err
	}
	t.Time = newTime
	return nil
}
