package models

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	ID int
	name              string
	age               int
	profesion         string
	slepping_time     time.Time
	average_work_time time.Time
	hourly_wage float64
	balance float64
}

func NewUser(id int, name string, age int, profesion string, slepping time.Time, work_time time.Time) *User {
	return &User{
		ID : id,
		name:              name,
		age:               age,
		profesion:         profesion,
		slepping_time:     slepping,
		average_work_time: work_time,
	}
}

func (u *User)Work(ctx context.Context,storage *Storage){
	startTime:=time.Now()
	fmt.Println("я начал работать!")
	time.Sleep(time.Duration(t.Hour())/60/4)
	endTime:=time.Now()
	daily_wage:=u.hourly_wage*float64(t.Hour())
	u.balance+=daily_wage
	NewWorkLog(*u,startTime,endTime)
	fmt.Printf("я отработал: %s\n",t.Format("2006-01-02 15:04:05"))
	fmt.Println("я заработал:",daily_wage)
}



