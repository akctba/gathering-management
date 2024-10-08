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

	// RSVP routes (public, but with gathering and invitee ID validation)
	r.POST("/rsvp/:gatheringID/:inviteeID", handlers.HandleRSVP)
	r.GET("/rsvp/:gatheringID/:inviteeID", handlers.GetRSVPStatus)

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

		// Invitee routes
		protected.POST("/gatherings/:gatheringID/invitees", handlers.CreateInvitee)
		protected.GET("/gatherings/:gatheringID/invitees", handlers.ListInvitees)
		protected.GET("/invitees/:id", handlers.GetInvitee)
		protected.PUT("/invitees/:id", handlers.UpdateInvitee)
		protected.DELETE("/invitees/:id", handlers.DeleteInvitee)

		// Food Plate routes
		protected.POST("/gatherings/:gatheringID/food-plates", handlers.CreateFoodPlate)
		protected.GET("/gatherings/:gatheringID/food-plates", handlers.ListFoodPlates)
		protected.GET("/food-plates/:id", handlers.GetFoodPlate)
		protected.PUT("/food-plates/:id", handlers.UpdateFoodPlate)
		protected.DELETE("/food-plates/:id", handlers.DeleteFoodPlate)

		// Beverage routes
		protected.POST("/gatherings/:gatheringID/beverages", handlers.CreateBeverage)
		protected.GET("/gatherings/:gatheringID/beverages", handlers.ListBeverages)
		protected.GET("/beverages/:id", handlers.GetBeverage)
		protected.PUT("/beverages/:id", handlers.UpdateBeverage)
		protected.DELETE("/beverages/:id", handlers.DeleteBeverage)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
