package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, con *pgx.Conn) error {
sqlquery:=`
CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	name VARCHAR(20) NOT NULL,
	age INTEGER,
	profession VARCHAR(20),
	sleeping_time TIMESTAMP,
	average_work_time TIMESTAMP
);
`
_,err := con.Exec(ctx,sqlquery)
return err
}
func CreateTableWorkLog(ctx context.Context, con *pgx.Conn)error{
	sqlquery:=`
CREATE TABLE IF NOT EXISTS work_log(
	id SERIAL PRIMARY KEY,
	user_id INTEGER REFERENCES user(id) ON DELETE CASCADE,
	start_time TIMESTAMP,
	end_time TIMESTAMP,
	comment TEXT
);
`
_,err := con.Exec(ctx,sqlquery)
return err
}
