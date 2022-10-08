package user

import (
	"fmt"

	"github.com/Hackathon-for-FUN-TeamA/backend/internal/db"
)

func Login(username, password string) (string, error) {
	dbmap, err := db.InitDb()
	if err != nil {
		return "", err
	}
	defer dbmap.Db.Close()

	// username, passwordがすでにdbにあるかチェック
	var result User
	err = dbmap.SelectOne(&result, "SELECT * FROM users WHERE username=? AND password=?", username, password)
	if err != nil {
		return "", err
	}

	// resultが空だったらerrorを返す
	if result.Username == "" && result.Password == "" {
		return "", fmt.Errorf("This user doesn't exist.")
	}

	return result.Token, nil
}
