package models

import (
	"time"
)

type User struct {
	ID int
	name              string
	age               int
	profesion         string
	slepping_time     time.Time
	average_work_time time.Time
	HourlyWage float64
	Balance float64
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

func (u *User)Work(hours float64) (earned float64){
	earned = hours * u.HourlyWage
	u.Balance += earned
	return earned
}



