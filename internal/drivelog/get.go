package drivelog

func Get(userId int, date string) (DriveLog, error) {
	dbmap, err := initDb()
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
