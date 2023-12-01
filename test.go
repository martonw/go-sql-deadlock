package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	var uuid sql.RawBytes
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	dbuser := os.Getenv("POSTGRES_USER")
	dbpass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	db, err := sql.Open("pgx", "postgres://"+dbuser+":"+dbpass+"@postgres:5432/"+dbname+"?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, "select id /* id is an uuid */ from test")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&uuid)
		fmt.Printf("uuid read: %v\n", uuid)
	}
	// cancel the context
	ctxCancel()

	fmt.Println("sleep begin")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("sleep end")
	// call rows.Err()
	fmt.Printf("err %v\n", rows.Err())
	fmt.Println("successful exit")

}
