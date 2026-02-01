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
	INSERT INTO users (name, age, profession, hourly_wage, balance)
	VALUES ($1,$2,$3,$4,$5) RETURNING id
	`
	row := pgUst.con.QueryRow(ctx, sqlQuery, u.Name, u.Age, u.Profession, u.HourlyWage, u.Balance)
	if err := row.Scan(&u.ID); err != nil {
		return err
	}
	return nil
}

func (pgUst *PgUserStore) AddEarnings(ctx context.Context, balance float64, id int)error {
	sqlQuery := `
	UPDATE users SET balance = balance + $1 WHERE id = $2
	`
	_,err:=pgUst.con.Exec(ctx, sqlQuery, balance, id)
	return err
}
