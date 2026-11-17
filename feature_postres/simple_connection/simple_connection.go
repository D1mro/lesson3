package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()
	conn,_:=pgx.Connect(ctx,"postgres://postgres:4711@localhost:5432/postgres")
	if err:=conn.Ping(ctx); err!=nil {
		panic(err)
	}
	fmt.Println("Подключение произошло успешно")
}


