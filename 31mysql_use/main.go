package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn_sql"
	db, err := sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		fmt.Println(recover())
		db.Close()
	}()
	err = db.Ping()
	if err != nil {
		fmt.Println("dsn",err)
		panic(1)
	}
}
