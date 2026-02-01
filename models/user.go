package models

import (
	"fmt"
	"math/rand"
	"strconv"
)

type User struct {
	Name              string
	Age               int
	Profession        string
	Slepping_time     int
	Average_work_time int
	HourlyWage        float64
	Balance           float64
}

func newUser(id int) *User {
	return &User{
		Name:              "person" + strconv.Itoa(id),
		Age:               rand.Intn(82) + 18,
		Profession:        "profession" + strconv.Itoa(id),
		Slepping_time:     rand.Intn(10) + 2,
		Average_work_time: rand.Intn(10) + 2,
		HourlyWage:        float64(rand.Intn(41) + 10),
	}
}

func (u *User) Work() (earned float64) {
	earned = float64(u.Average_work_time) * u.HourlyWage
	u.Balance += earned
	fmt.Println("Я заработал деньги:", u.Balance)
	return earned
}
