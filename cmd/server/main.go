package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// sample
	r.GET("/ping", Ping)

	// sample
	r.POST("/throw", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
	})

	// ユーザ作成
	r.POST("/user/create", UserCreate)
	// ログイン
	r.POST("/login", Login)
	// ドライブ時のログを保存
	r.POST("/drivelog", PostDriveLog)
	// ドライブ時のログを保存
	r.GET("/drivelog", GetDriveLog)

	r.Run() // 127.0.0.0:8000でサーバを建てる
}
