package models

import "context"

type UserStore interface {
	CreateUser(ctx context.Context, u *User) error
	AddEarnings(ctx context.Context, userID int, earned float64) error
}
