package handlers

import (
	"net/http"
	"strconv"

	"github.com/akctba/gathering-management/configs"
	"github.com/akctba/gathering-management/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateFoodPlate handles the creation of a new food plate
func CreateFoodPlate(c *gin.Context) {
	var input models.FoodPlate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gatheringID := c.Param("gatheringID")
	gatheringIDUint, err := strconv.ParseUint(gatheringID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gathering ID"})
		return
	}
	input.GatheringID = uint(gatheringIDUint)

	if err := configs.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create food plate"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// GetFoodPlate retrieves a single food plate by ID
func GetFoodPlate(c *gin.Context) {
	id := c.Param("id")

	var foodPlate models.FoodPlate
	if err := configs.DB.First(&foodPlate, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Food plate not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodPlate})
}

// ListFoodPlates retrieves all food plates for a specific gathering
func ListFoodPlates(c *gin.Context) {
	gatheringID := c.Param("gatheringID")

	var foodPlates []models.FoodPlate
	if err := configs.DB.Where("gathering_id = ?", gatheringID).Find(&foodPlates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve food plates"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foodPlates})
}

// UpdateFoodPlate updates a food plate
func UpdateFoodPlate(c *gin.Context) {
	id := c.Param("id")

	var foodPlate models.FoodPlate
	if err := configs.DB.First(&foodPlate, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Food plate not found"})
		return
	}

	var input models.FoodPlate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configs.DB.Model(&foodPlate).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": foodPlate})
}

// DeleteFoodPlate deletes a food plate
func DeleteFoodPlate(c *gin.Context) {
	id := c.Param("id")

	var foodPlate models.FoodPlate
	if err := configs.DB.First(&foodPlate, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Food plate not found"})
		return
	}

	configs.DB.Delete(&foodPlate)

	c.JSON(http.StatusOK, gin.H{"data": "Food plate deleted successfully"})
}
