package main

import (
	"log"

	"github.com/akctba/gathering-management/configs"
	"github.com/akctba/gathering-management/internal/handlers"
	"github.com/akctba/gathering-management/internal/middleware"
	"github.com/akctba/gathering-management/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	configs.ConnectDatabase()

	// Auto Migrate
	configs.DB.AutoMigrate(&models.User{}, &models.Gathering{}, &models.Invitee{}, &models.FoodPlate{}, &models.Beverage{})

	r := gin.Default()

	// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Gathering routes
		protected.POST("/gatherings", handlers.CreateGathering)
		protected.GET("/gatherings", handlers.ListGatherings)
		protected.GET("/gatherings/:id", handlers.GetGathering)
		protected.PUT("/gatherings/:id", handlers.UpdateGathering)
		protected.DELETE("/gatherings/:id", handlers.DeleteGathering)

		// Placeholder for future routes
		protected.GET("/ping", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			c.JSON(200, gin.H{
				"message": "pong",
				"userID":  userID,
			})
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
