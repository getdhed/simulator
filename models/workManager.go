package models

import "context"

type WorkManager struct {
	store UserStore
}

func NewWorkManager(store UserStore) *WorkManager {
	return &WorkManager{store: store}
}
func (wm *WorkManager) StartShift(ctx context.Context, mp *Users) error {
	for id, user := range mp.users {
		earned := user.Work()
		if err := wm.store.AddEarnings(ctx, id, earned); err != nil {
			return err
		}
	}
	return nil
}
