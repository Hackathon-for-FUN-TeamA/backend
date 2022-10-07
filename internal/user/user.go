package user

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"gopkg.in/gorp.v1"
)

type User struct {
	Id       int
	Token    string
	Username string
	Password string
}

func CreateUser(username, password string) (string, error) {
	dbmap, err := initDb()
	if err != nil {
		return "", err
	}
	defer dbmap.Db.Close()

	// token生成
	uuidobj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(User{}, "users")

	// insert
	userData := &User{
		Token:    uuidobj.String(),
		Username: username,
		Password: password,
	}
	err = dbmap.Insert(userData)
	if err != nil {
		return "", err
	}

	return uuidobj.String(), nil
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
