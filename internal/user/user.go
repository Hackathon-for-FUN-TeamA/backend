package user

import (
	"github.com/Hackathon-for-FUN-TeamA/backend/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func CreateUser(username, password string) (string, error) {
	dbmap, err := db.InitDb()
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
