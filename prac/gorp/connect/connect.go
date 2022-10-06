package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	path := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=true",
		"root", "password", "map")
	db, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Printf("open error: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("ping error: %v\n", err)
	}

	fmt.Println("Success!")
}
