package models

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PgUserStore struct {
	con *pgx.Conn
}

func NewPgUserStore(con *pgx.Conn) *PgUserStore {
	return &PgUserStore{con: con}
}

func (pgUst *PgUserStore) CreateUser(ctx context.Context, u *User) error {
	sqlQuery := `
	INSERT INTO users (id, name, age, profession, sleeping_time, average_work_time, hourly_wage, balance)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	ON CONFLICT (id) DO NOTHING
	`
	_, err := pgUst.con.Exec(ctx, sqlQuery,
		u.ID,
		u.Name,
		u.Age,
		u.Profession,
		u.Sleeping_time,     // или как у тебя поле называется
		u.Average_work_time, // или как у тебя поле называется
		u.HourlyWage,
		u.Balance,
	)

	return err
}

func (pgUst *PgUserStore) AddEarnings(ctx context.Context, id int, balance float64) error {
	sqlQuery := `
	UPDATE users SET balance = balance + $1 WHERE id = $2
	`
	_, err := pgUst.con.Exec(ctx, sqlQuery, balance, id)
	return err
}

func (pgUst *PgUserStore) ListUsers(ctx context.Context) ([]*User, error) {
	sqlQuery := `
 	SELECT id,name, age, profession, sleeping_time, average_work_time, hourly_wage, balance FROM users
 	`

	rows, err := pgUst.con.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		u := &User{}
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Age,
			&u.Profession,
			&u.Sleeping_time,
			&u.Average_work_time,
			&u.HourlyWage,
			&u.Balance,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}
