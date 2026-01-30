package models

type Users struct {
	users map[int]*User
}

func NewUsers() *Users {
	return &Users{
		users: make(map[int]*User),
	}
}
func (u *Users) Add(user *User) {
	u.users[user.ID] = user
}
func (u *Users) Get(id int) (*User, bool) {
	user, ok := u.users[id]
	return user, ok
}
