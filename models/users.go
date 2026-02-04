package models

import (
	"context"
	"fmt"
)

type Users struct {
	users  map[int]*User
	nextID int
}

func NewUsers() *Users {
	return &Users{
		users:  make(map[int]*User),
		nextID: 1,
	}
}
func UsersFromDB(sl []*User) *Users {
	u := &Users{
		users:  make(map[int]*User),
		nextID: 1,
	}
	for _, user := range sl {
		u.users[user.ID] = user
		if user.ID >= u.nextID {
			u.nextID = user.ID + 1
		}
	}
	return u
}

func (u *Users) Add() *User {
	id := u.nextID
	user := newUser(id)
	u.users[id] = user
	u.nextID++
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
			"ID: %d\n| Name: %s\n| Age: %d\n| Profession: %s\n| Balance: %f\n",
			id,
			user.Name,
			user.Age,
			user.Profession,
			user.Balance,
		)
	}
}

func (u *Users) ForEach(fn func(*User)) {
	for _, u := range u.users {
		fn(u)
	}
}

func (u *Users) CreateAll(ctx context.Context, store UserStore) error {
	for _, user := range u.users {
		if err := store.CreateUser(ctx, user); err != nil {
			return err
		}
	}
	return nil
}
