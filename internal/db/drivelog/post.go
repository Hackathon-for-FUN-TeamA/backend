package drivelog

import (
	"github.com/Hackathon-for-FUN-TeamA/backend/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func Post(userId int, date string, speed, acceleration, latitude, longtitude float64) error {
	dbmap, err := db.InitDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(DriveLog{}, "drivelogs")

	// insert
	userData := &DriveLog{
		User_Id:      userId,
		Date:         date,
		Speed:        speed,
		Acceleration: acceleration,
		Latitude:     latitude,
		Longtitude:   longtitude,
	}
	err = dbmap.Insert(userData)
	if err != nil {
		return err
	}

	return nil
}
