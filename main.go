package main

import (
	"context"
	"fmt"
	"lesson3/feature_postgres/simple_connection"
	"lesson3/feature_postgres/simple_sql"
)

func main() {
	ctx := context.Background()
	conn, _ := simple_connection.CreateConnection(ctx)

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("Таблица создана успешно!")

	// if err := simple_sql.InsertRow(ctx,
	// 	conn,
	// 	"Обед",
	// 	"soup",
	// 	false,
	// 	time.Now(),
	// 	); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.UpdateRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := simple_sql.DeleteRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	tasks, _ := simple_sql.SelectRows(ctx, conn)
	fmt.Println(tasks)

	fmt.Println("Succeed!")

}
