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
