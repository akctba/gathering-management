package handlers

import (
	"net/http"

	"github.com/akctba/gathering-management/configs"
	"github.com/akctba/gathering-management/internal/models"
	"github.com/gin-gonic/gin"
)

// HandleRSVP processes the RSVP response from an invitee
func HandleRSVP(c *gin.Context) {
	inviteeID := c.Param("inviteeID")
	gatheringID := c.Param("gatheringID")

	var input struct {
		Status string `json:"status" binding:"required,oneof=accepted declined"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var invitee models.Invitee
	if err := configs.DB.Where("id = ? AND gathering_id = ?", inviteeID, gatheringID).First(&invitee).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invitee not found or does not belong to this gathering"})
		return
	}

	invitee.RSVPStatus = input.Status
	if err := configs.DB.Save(&invitee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update RSVP status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "RSVP status updated successfully", "status": input.Status})
}

// GetRSVPStatus retrieves the current RSVP status for an invitee
func GetRSVPStatus(c *gin.Context) {
	inviteeID := c.Param("inviteeID")
	gatheringID := c.Param("gatheringID")

	var invitee models.Invitee
	if err := configs.DB.Where("id = ? AND gathering_id = ?", inviteeID, gatheringID).First(&invitee).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invitee not found or does not belong to this gathering"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": invitee.RSVPStatus})
}
