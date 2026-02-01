package models

type WorkManager struct {
	users *Users
}

func NewWorkManager() *WorkManager {
	return &WorkManager{
		users: NewUsers(),
	}
}
func (wm *WorkManager) StartShift(mp *Users) {
	for _, user := range mp.users {
		user.Work()
	}
}
