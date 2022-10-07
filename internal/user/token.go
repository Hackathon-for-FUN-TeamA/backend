package user

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetUserByToken(token string) (int, error) {
	dbmap, err := initDb()
	if err != nil {
		return -1, err
	}
	defer dbmap.Db.Close()

	// tokenからuserを探す
	var result User
	err = dbmap.SelectOne(&result, "SELECT * FROM users WHERE token=?", token)
	if err != nil {
		return -1, err
	}

	// resultが空だったらerrorを返す
	if result.Id <= 0 {
		return -1, fmt.Errorf("This user doesn't exist.")
	}

	return result.Id, nil
}
