package main

import (
	"github.com/Hackathon-for-FUN-TeamA/backend/internal/drivelog"
	"github.com/Hackathon-for-FUN-TeamA/backend/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// sample
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ユーザ作成
	r.POST("/user/create", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.JSON(400, gin.H{
				"message": "query string is null",
			})
		}

		token, err := user.CreateUser(username, password)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{
				"token": token,
			})
		}

	})

	// ログイン
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.JSON(400, gin.H{
				"message": "query string is null",
			})
		}

		token, err := user.Login(username, password)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{
				"token": token,
			})
		}
	})

	// ドライブ時のログを保存
	r.POST("/drivelog", func(c *gin.Context) {
		// param取得
		token := c.PostForm("token")
		date := c.PostForm("date")                               // 日時・時間
		speed := c.GetFloat64(c.PostForm("speed"))               // 速度
		acceleration := c.GetFloat64(c.PostForm("acceleration")) // 加速度
		latitude := c.GetFloat64(c.PostForm("latitude"))         // 緯度
		longtude := c.GetFloat64(c.PostForm("longtude"))         // 経度

		userId, err := user.GetUserByToken(token)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
				"token":   token,
				"user_id": userId,
			})
		}

		// tokenが無効な場合
		if userId == -1 {
			c.JSON(401, gin.H{
				"message": "Not Credential",
			})
		} else {
			err := drivelog.Post(userId, date, speed, acceleration, latitude, longtude)
			// TODO: 入れ子になってて可読性が下がりそう。どうにかしたい
			if err != nil {
				c.JSON(402, gin.H{
					"message": err,
				})
			} else {
				c.JSON(200, gin.H{})
			}
		}

	})

	// ドライブ時のログを保存
	r.GET("/drivelog", func(c *gin.Context) {
		// param取得
		token := c.PostForm("token")
		date := c.PostForm("date") // 日時・時間

		userId, err := user.GetUserByToken(token)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
		}

		// tokenが無効な場合
		if userId == -1 {
			c.JSON(400, gin.H{
				"message": "Not Credential",
			})
		} else {
			drivelog, err := drivelog.Get(userId, date)
			// TODO: 入れ子になってて可読性が下がりそう。どうにかしたい
			if err != nil {
				c.JSON(400, gin.H{
					"message": "Not Credential",
				})
			} else {
				c.JSON(200, gin.H{
					"speed":        drivelog.Speed,
					"acceleration": drivelog.Acceleration,
					"latitude":     drivelog.Latitude,
					"longtude":     drivelog.Longtitude,
				})
			}
		}

	})

	r.Run() // 127.0.0.0:8000でサーバを建てる
}
