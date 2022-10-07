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

	r.POST("/user/create", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.JSON(400, gin.H{
				"message": "query string is null",
			})
		}

		err := user.CreateUser(username, password)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Created",
			})
		}

	})

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

	r.Run() // 127.0.0.0:8000でサーバを建てる
}
