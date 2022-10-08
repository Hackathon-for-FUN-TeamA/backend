package drivelog

import "github.com/Hackathon-for-FUN-TeamA/backend/internal/db"

func Get(userId int, date string) (DriveLog, error) {
	dbmap, err := db.InitDb()
	if err != nil {
		return DriveLog{}, err
	}
	defer dbmap.Db.Close()

	// username, passwordがすでにdbにあるかチェック
	var result DriveLog
	err = dbmap.SelectOne(&result, "SELECT * FROM drivelogs WHERE user_id=? AND date=?", userId, date)
	if err != nil {
		return DriveLog{}, err
	}

	return result, nil
}
