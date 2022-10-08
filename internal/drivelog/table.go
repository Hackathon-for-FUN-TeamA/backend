package drivelog

type DriveLog struct {
	LogId        int
	UserId       int
	Date         string
	Speed        float64
	Acceleration float64
	Latitude     float64
	Longtitude   float64
}
