package drivelog

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

type DriveLog struct {
	LogId        int
	UserId       int
	Date         string
	Speed        float64
	Acceleration float64
	Latitude     float64
	Longtitude   float64
}

func Post(userId int, date string, speed, acceleration, latitude, longtitude float64) error {
	dbmap, err := initDb()
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	// dbmapにテーブル名登録
	dbmap.AddTableWithName(DriveLog{}, "drivelogs")

	// insert
	userData := &DriveLog{
		UserId:       userId,
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

func initDb() (*gorp.DbMap, error) {
	// db接続
	path := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=true",
		"root",    // username
		"passwd",  // password
		"asayake", // database
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
