package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	Pool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{Pool: pool}
}

func (s *Storage) CreateUser(ctx context.Context, u *User) error {
	_, err := s.Pool.Exec(ctx,
		`INSERT INTO users (id, name, age, profession, hourly_wage, balance)
		 VALUES ($1,$2,$3,$4,$5,$6)`,
		u.ID, u.Name, u.Age, u.Profession, u.HourlyWage, u.Balance,
	)
	return err
}

func (s *Storage) InsertWorkLog(ctx context.Context, log *WorkLog) error {
	_, err := s.Pool.Exec(ctx,
		`INSERT INTO work_logs (user_id, start_time, end_time)
		 VALUES ($1,$2,$3)`,
		log.UserID, log.StartTime, log.EndTime,
	)
	return err
}