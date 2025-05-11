package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/v1/greeting", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/v1/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ini adlaah hands-on deployment VPS",
		})
	})

	r.Run(":8000")
}
