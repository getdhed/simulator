package models

import "time"

type WorkLog struct {
	UserID    int
	StartTime time.Time
	EndTime   time.Time
	Comment   string
}

func NewWorkLog(u User, st time.Time, end time.Time) *WorkLog {
	return &WorkLog{
		StartTime: st,
		EndTime:   end,
	}
}
