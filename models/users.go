package models

import "fmt"

type Users struct {
	users  map[int]*User
	lastID int
}

func NewUsers() *Users {
	return &Users{
		users: make(map[int]*User),
	}
}

func (u *Users) Add() *User {
	u.lastID++
	user := newUser(u.lastID)
	u.users[u.lastID] = user
	return user
}

func (u *Users) Get(id int) (*User, bool) {
	user, ok := u.users[id]
	return user, ok
}

func (u *Users) InsertMap(n int) {
	for i := 0; i < n; i++ {
		u.Add()
	}
}

func (u *Users) PrintMap() {
	for id, user := range u.users {
		fmt.Printf(
			"ID: %d\n| Name: %s\n| Age: %d\n| Profession: %s\n| Balance: %s\n",
			id,
			user.Name,
			user.Age,
			user.Profession,
			user.Balance,
		)
	}
}
