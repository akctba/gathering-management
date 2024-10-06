package handlers

import (
	"net/http"

	"github.com/akctba/gathering-management/configs"
	"github.com/akctba/gathering-management/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateGathering handles the creation of a new gathering
func CreateGathering(c *gin.Context) {
	var input models.Gathering
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	input.CreatorID = userID.(uint)

	if err := configs.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create gathering"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// GetGathering retrieves a single gathering by ID
func GetGathering(c *gin.Context) {
	id := c.Param("id")

	var gathering models.Gathering
	if err := configs.DB.Preload("Creator").Preload("Invitees").First(&gathering, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gathering not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gathering})
}

// ListGatherings retrieves all gatherings for the authenticated user
func ListGatherings(c *gin.Context) {
	userID, _ := c.Get("userID")

	var gatherings []models.Gathering
	if err := configs.DB.Where("creator_id = ?", userID).Find(&gatherings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve gatherings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gatherings})
}

// UpdateGathering updates a gathering
func UpdateGathering(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")

	var gathering models.Gathering
	if err := configs.DB.First(&gathering, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gathering not found"})
		return
	}

	if gathering.CreatorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this gathering"})
		return
	}

	var input models.Gathering
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configs.DB.Model(&gathering).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": gathering})
}

// DeleteGathering deletes a gathering
func DeleteGathering(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")

	var gathering models.Gathering
	if err := configs.DB.First(&gathering, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gathering not found"})
		return
	}

	if gathering.CreatorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this gathering"})
		return
	}

	configs.DB.Delete(&gathering)

	c.JSON(http.StatusOK, gin.H{"data": "Gathering deleted successfully"})
}
