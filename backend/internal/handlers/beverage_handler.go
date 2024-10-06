package handlers

import (
	"net/http"

	"strconv"

	"github.com/akctba/gathering-management/configs"
	"github.com/akctba/gathering-management/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateBeverage handles the creation of a new beverage
func CreateBeverage(c *gin.Context) {
	var input models.Beverage
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create beverage"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// GetBeverage retrieves a single beverage by ID
func GetBeverage(c *gin.Context) {
	id := c.Param("id")

	var beverage models.Beverage
	if err := configs.DB.First(&beverage, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Beverage not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": beverage})
}

// ListBeverages retrieves all beverages for a specific gathering
func ListBeverages(c *gin.Context) {
	gatheringID := c.Param("gatheringID")

	var beverages []models.Beverage
	if err := configs.DB.Where("gathering_id = ?", gatheringID).Find(&beverages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve beverages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": beverages})
}

// UpdateBeverage updates a beverage
func UpdateBeverage(c *gin.Context) {
	id := c.Param("id")

	var beverage models.Beverage
	if err := configs.DB.First(&beverage, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Beverage not found"})
		return
	}

	var input models.Beverage
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configs.DB.Model(&beverage).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": beverage})
}

// DeleteBeverage deletes a beverage
func DeleteBeverage(c *gin.Context) {
	id := c.Param("id")

	var beverage models.Beverage
	if err := configs.DB.First(&beverage, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Beverage not found"})
		return
	}

	configs.DB.Delete(&beverage)

	c.JSON(http.StatusOK, gin.H{"data": "Beverage deleted successfully"})
}
