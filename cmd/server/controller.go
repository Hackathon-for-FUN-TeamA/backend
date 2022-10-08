package main

import (
	"github.com/Hackathon-for-FUN-TeamA/backend/internal/db/drivelog"
	"github.com/Hackathon-for-FUN-TeamA/backend/internal/db/user"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserCreate(c *gin.Context) {
	// パラメータ取得
	type JsonRequest struct {
		UserName string `json: "username"`
		Password string `json: "password"`
	}

	var req JsonRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	token, err := user.CreateUser(req.UserName, req.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func Login(c *gin.Context) {
	// パラメータ取得
	type JsonRequest struct {
		UserName string `json: "username"`
		Password string `json: "password"`
	}

	var req JsonRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	token, err := user.Login(req.UserName, req.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func PostDriveLog(c *gin.Context) {
	// param取得
	type JsonRequest struct {
		// UserName string `json: "username"`
		Token        string  `json: "token"`
		Date         string  `json: "date"`         // 現在の日時
		Speed        float64 `json: "speed"`        // 速度
		Acceleration float64 `json: "acceleration"` // 加速度
		Latitude     float64 `json: "latitude"`     // 緯度
		Longtitude   float64 `json: "longtitude"`   // 経度
	}

	var req JsonRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	userId, err := user.GetUserByToken(req.Token)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	// tokenが無効な場合
	if userId == -1 {
		c.JSON(401, gin.H{
			"message": "UNAUTHORIZED",
		})
		return
	}

	err = drivelog.Post(userId, req.Date, req.Speed, req.Acceleration, req.Latitude, req.Longtitude)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{})

}

func GetDriveLog(c *gin.Context) {
	// param取得
	type JsonRequest struct {
		Token string `json: "token"`
		Date  string `json: "date"` // 現在の日時
	}
	var req JsonRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	userId, err := user.GetUserByToken(req.Token)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	// tokenが無効な場合
	if userId == -1 {
		c.JSON(401, gin.H{
			"message": "UNAUTHORIZED",
		})
		return
	}

	drivelog, err := drivelog.Get(userId, req.Date)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"speed":        drivelog.Speed,
		"acceleration": drivelog.Acceleration,
		"latitude":     drivelog.Latitude,
		"longtude":     drivelog.Longtitude,
	})

}
