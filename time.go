package goWallabag

import (
	"time"
)

//TimeFormat time format use by API
const TimeFormat = "2006-01-02T15:04:05-0700"

//Time struct use for specific JSON Unmarsahl cause of format
type Time struct {
	time.Time
}

//UnmarshalJSON func use for parsing time with TimeFormat
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
