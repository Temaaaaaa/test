package main

import (
	"context"
	"db/feature_postgres/simple_connection"
	"db/feature_postgres/simple_sql"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	//if err := simple_sql.CreateTable(ctx, conn); err != nil {
	//	panic(err)
	//}
	//if err := simple_sql.InsertRow(ctx, conn, "Обед", "Покушать", false, time.Now()); err != nil {
	//	panic(err)
	//}
	//fmt.Println("Successfully")
	//if err := simple_sql.Update(ctx, conn); err != nil {
	//	panic(err)
	//}

	if err := simple_sql.UpdateTask(ctx, conn, simple_sql.TaskModel{
		ID:          1,
		Title:       "Купить негра",
		Description: "Секс сиськи",
		Completed:   true,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}); err != nil {
		// Теперь мы точно узнаем, если что-то пойдет не так!
		panic(fmt.Errorf("не удалось обновить задачу: %w", err))
	}
	tasks, err := simple_sql.Select(ctx, conn)
	if err != nil {
		panic(err)
	}
	fmt.Println(tasks)

	fmt.Println("Successfully")

}
