package models

import "context"

type WorkManager struct {
	store UserStore
}

func NewWorkManager(store UserStore) *WorkManager {
	return &WorkManager{store: store}
}

func (wm *WorkManager) SimulateDay(ctx context.Context, day int, users *Users) (DayResult, error) {
	res := DayResult{
		Day:     day,
		PerUser: make(map[int]float64, len(users.users)),
	}

	for id, u := range users.users {
		earned := u.Work()

		// 1) сохранить в БД баланс
		if err := wm.store.AddEarnings(ctx, id, earned); err != nil {
			return DayResult{}, err
		}

		// 2) собрать статистику дня
		res.Workers++
		res.TotalEarned += earned
		res.PerUser[id] = earned
	}

	return res, nil
}
