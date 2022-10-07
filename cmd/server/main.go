package main

import (
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
				"error": "query string is null",
			})
		}

		token, err := user.CreateUser(username, password)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
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

		// TODO: authとか作って認証できるとよい
		err := user.Login(username, password)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Login",
			})
		}
	})

	// ドライブ時のログを保存
	r.POST("/drivelog", func(c *gin.Context) {

	})

	r.Run() // 127.0.0.0:8000でサーバを建てる
}
