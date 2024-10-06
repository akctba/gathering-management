package handlers

import (
	"net/http"
	"strconv"

	"github.com/akctba/gathering-management/configs"
	"github.com/akctba/gathering-management/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateInvitee handles the creation of a new invitee
func CreateInvitee(c *gin.Context) {
	var input models.Invitee
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create invitee"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// GetInvitee retrieves a single invitee by ID
func GetInvitee(c *gin.Context) {
	id := c.Param("id")

	var invitee models.Invitee
	if err := configs.DB.First(&invitee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invitee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invitee})
}

// ListInvitees retrieves all invitees for a specific gathering
func ListInvitees(c *gin.Context) {
	gatheringID := c.Param("gatheringID")

	var invitees []models.Invitee
	if err := configs.DB.Where("gathering_id = ?", gatheringID).Find(&invitees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve invitees"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invitees})
}

// UpdateInvitee updates an invitee
func UpdateInvitee(c *gin.Context) {
	id := c.Param("id")

	var invitee models.Invitee
	if err := configs.DB.First(&invitee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invitee not found"})
		return
	}

	var input models.Invitee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configs.DB.Model(&invitee).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": invitee})
}

// DeleteInvitee deletes an invitee
func DeleteInvitee(c *gin.Context) {
	id := c.Param("id")

	var invitee models.Invitee
	if err := configs.DB.First(&invitee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invitee not found"})
		return
	}

	configs.DB.Delete(&invitee)

	c.JSON(http.StatusOK, gin.H{"data": "Invitee deleted successfully"})
}
