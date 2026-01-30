package models

import "fmt"

type User struct {
	ID                int
	Name              string
	Age               int
	Profession        string
	Slepping_time     int
	Average_work_time int
	HourlyWage        float64
	Balance           float64
}

func NewUser(id int, name string, age int, profession string, slepping int, work_time int) *User {
	return &User{
		ID:                id,
		Name:              name,
		Age:               age,
		Profession:        profession,
		Slepping_time:     slepping,
		Average_work_time: work_time,
	}
}

func (u *User) Work() (earned float64) {
	earned = float64(u.Average_work_time) * u.HourlyWage
	u.Balance += earned
	fmt.Println("Я заработал деньги:", u.Balance)
	return earned
}
