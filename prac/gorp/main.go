package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

type User struct {
	Id       int
	Username string
	Password string
}

func main() {
	dbmap, err := initDb()
	if err != nil {
		log.Fatal("error initDb", err)
	}
	defer dbmap.Db.Close()

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(User{}, "user")

	// select
	// _, err := dbmap.Select(&users)
	// insertする
	user1 := &User{
		Id:       100,
		Username: "Alice",
		Password: "passw",
	}
	err = dbmap.Insert(user1)
	if err != nil {
		log.Fatal("error insert", err)
	}
}

func initDb() (*gorp.DbMap, error) {
	// db接続
	path := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=true",
		"root", "password", "asayake",
	)
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}

	// dbをdbmapに
	dbmap := &gorp.DbMap{
		Db:      db,
		Dialect: gorp.MySQLDialect{},
	}

	return dbmap, nil
}
