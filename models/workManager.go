package models

import "context"

type WorkManager struct {
	users *Users
}

func NewWorkManager() *WorkManager {
	return &WorkManager{
		users: NewUsers(),
	}
}
func (wm *WorkManager) StartShift(ctx context.Context,mp *Users, sqlStore *PgUserStore) {

	for id, user := range mp.users {

		earned := user.Work()
		sqlStore.AddEarnings(ctx,earned,mp.users[id].ID)

	}
}
