package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/akctba/gathering-management/configs"
	"github.com/akctba/gathering-management/internal/models"
	"github.com/akctba/gathering-management/internal/services"
	"github.com/gin-gonic/gin"
)

var emailService = services.NewEmailService()

// CreateInvitee handles the creation of a new invitee
func CreateInvitee(c *gin.Context) {
	var input models.Invitee
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gatheringID, err := strconv.ParseUint(c.Param("gatheringID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gathering ID"})
		return
	}
	input.GatheringID = uint(gatheringID)

	if err := configs.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create invitee"})
		return
	}

	// Fetch the gathering to get its name
	var gathering models.Gathering
	if err := configs.DB.First(&gathering, gatheringID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch gathering details"})
		return
	}

	// Generate invite link (TODO: create a separate function for this)
	inviteLink := fmt.Sprintf("%s/rsvp/%d/%d", os.Getenv("APP_URL"), gatheringID, input.ID)

	// Send invitation email
	// TODO: handle errors
	// TODO: move this to a service
	// TODO: add a boolean flag to the invitee model to track if the email was sent
	// TODO: add a goolean attribute on request body to send email or not
	if err := emailService.SendInvitation(input.Email, gathering.Name, inviteLink); err != nil {
		// Log the error, but don't return it to the client
		log.Printf("Failed to send invitation email: %v", err)
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
