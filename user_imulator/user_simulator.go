package userimulator

import "time"

type User struct {
	id                int
	name              string
	age               int
	profesion         string
	slepping_time     time.Time
	average_work_time time.Time
}

func NewUser(id int, name string, age int, profesion string, slepping time.Time, work_time time.Time) *User {
	return &User{
		id:                id,
		name:              name,
		age:               age,
		profesion:         profesion,
		slepping_time:     slepping,
		average_work_time: work_time,
	}
}
