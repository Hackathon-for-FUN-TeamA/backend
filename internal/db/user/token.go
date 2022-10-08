package user

import (
	"github.com/Hackathon-for-FUN-TeamA/backend/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

// tokenからuserを取得
func GetUserByToken(token string) (int, error) {
	dbmap, err := db.InitDb()
	if err != nil {
		return -1, err
	}
	defer dbmap.Db.Close()

	// tokenからuserを探す
	var result User
	err = dbmap.SelectOne(&result, "SELECT * FROM `users` WHERE `token`=?", token)
	if err != nil {
		return -1, err
	}

	return result.Id, nil
}
