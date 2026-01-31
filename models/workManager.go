package models

type WorkManager struct {
	users *Users
}

func (wm *WorkManager) NewWorkManager() *WorkManager {
	return &WorkManager{}
}
