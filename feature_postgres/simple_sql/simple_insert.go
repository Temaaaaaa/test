package simple_sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

// Добавляем title, description, completed и createdAt в аргументы функции
func InsertRow(ctx context.Context, conn *pgx.Conn, title, description string, completed bool, createdAt time.Time) error {
	Query := `
    INSERT INTO tasks (title, description, completed, created_at)
    VALUES ($1, $2, $3, $4); -- Убрали одинарные кавычки
    `
	// Передаем аргументы в conn.Exec
	_, err := conn.Exec(ctx, Query, title, description, completed, createdAt)
	return err
}
