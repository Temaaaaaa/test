package simple_sql

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func UpdateTask(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	query := `
    UPDATE tasks
    SET title=$1,
        description=$2,
        completed=$3,
        created_at=$4,
        completed_at=$5
    WHERE id = $6;
    `
	_, err := conn.Exec(ctx, query,
		task.Title,       // $1
		task.Description, // $2
		task.Completed,   // $3
		task.CreatedAt,   // $4 - теперь соответствует created_at=$4
		task.CompletedAt, // $5 - теперь соответствует completed_at=$5
		task.ID,          // $6
	)
	return err
}
