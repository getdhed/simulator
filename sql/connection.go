package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connection(ctx context.Context, con string) error {

	_, err := pgx.Connect(ctx, con)
	return err
}
