package user

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

type User struct {
	Username string
	Password string
}

func CreateUser(username, password string) error {
	dbmap, err := initDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(User{}, "users")

	userData := &User{
		Username: username,
		Password: password,
	}
	err = dbmap.Insert(userData)
	if err != nil {
		return err
	}

	return nil
}

func initDb() (*gorp.DbMap, error) {
	// db接続
	path := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=true",
		"root",     // username
		"password", // password
		"asayake",  // database
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
