package main

import (
	"context"
	"go-deploy-vps/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// Initialize the database connection
	db := config.InitDatabase()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Println("Error finding database instance:", err)
		}
		sqlDB.Close()
	}()

	// Initialize Redis connection
	ctx := context.Background()
	redisClient := config.ConnectRedis(ctx)
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Println("Error closing Redis connection:", err)
		}
	}()

	r := gin.Default()

	r.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong!",
		})
	})

	r.GET("/v1/greeting", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Mentor!",
		})
	})

	r.GET("/v1/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ini adalah hands-on deployment VPS pada bootcamp dibimbing",
		})
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	log.Println("Server is running on port " + port)
	r.Run(":" + port)
}
