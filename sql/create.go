package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, con *pgx.Conn) error {
	sqlquery := `
CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	name VARCHAR(20) NOT NULL,
	age INTEGER,
	profession VARCHAR(20),
	sleeping_time INTEGER,
	average_work_time INTEGER,
	hourly_wage NUMERIC(10,2) NOT NULL,
	balance NUMERIC(12,2) NOT NULL DEFAULT 0
);
`
	_, err := con.Exec(ctx, sqlquery)
	return err
}
func CreateTableWorkLog(ctx context.Context, con *pgx.Conn) error {
	sqlquery := `
CREATE TABLE IF NOT EXISTS work_log(
	id SERIAL PRIMARY KEY,
	user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
	earned NUMERIC(10,2) NOT NULL,
	work_time INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT NOW()
);
`
	_, err := con.Exec(ctx, sqlquery)
	return err
}
