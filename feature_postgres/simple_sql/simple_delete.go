package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Delete(ctx context.Context, conn *pgx.Conn, taskID int) error {
	Query := `
	DELETE FROM tasks
	WHERE id = $1
	`
	_, err := conn.Exec(ctx, Query, taskID)
	return err

}
