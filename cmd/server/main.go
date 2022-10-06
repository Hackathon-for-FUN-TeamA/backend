package main

import (
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

	})

	r.Run() // 0.0.0.0:8000でサーバを建てる
}
