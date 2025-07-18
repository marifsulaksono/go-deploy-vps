package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

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
			"message": "Ini adalah hands-on deployment VPS dibimbing",
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	r.Run(":" + port)
}
