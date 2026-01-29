package userimulator

import "time"

type WorkLog struct {
	ID        int
	UserID    int
	StartTime time.Time
	EndTime   time.Time
	Comment   string
}
func (w *WorkLog)NewWorkLog(u User,st time.Time,end time.Time)*WorkLog{
	return &WorkLog{
		//ID: ?,
		UserID: u.id,
		StartTime: st,
EndTime: end,
	}
}