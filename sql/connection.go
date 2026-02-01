package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connection(ctx context.Context, con string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, con)
	return conn, err
}
