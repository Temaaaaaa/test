package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Select(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	query := `
    SELECT id, title, description, completed, created_at, completed_at
    FROM tasks
    WHERE completed = TRUE
    LIMIT 10
    OFFSET 0
    `
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks = make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		// ИСПРАВЛЕНО: createdAt идет ПЕРЕД completedAt, как и в SELECT
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.CompletedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
		//PrintTask(task)
	}

	return tasks, nil
}

func PrintTask(task TaskModel) {
	fmt.Println("id", task.ID)
	fmt.Println("title", task.Title)
	fmt.Println("descr", task.Description)
	fmt.Println("completed", task.Completed)
	fmt.Println("created at", task.CreatedAt)
	fmt.Println("completed at", task.CompletedAt)
}
