package goWallabag

import (
	"time"
)

const TimeFormat = "2006-01-02T15:04:05-0700"

type Time struct {
	time.Time
}

func (self *Time) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]

	t, err := time.Parse(TimeFormat, s)
	if err != nil {
		return err
	}
	self.Time = t
	return nil
}
