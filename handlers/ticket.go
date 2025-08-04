package handlers

import (
	"net/http"

	"delfi-scanner-api/db"
	"delfi-scanner-api/models"

	"github.com/gin-gonic/gin"
)

func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Define status as default - "Pending"
	ticket.TicketStatus = models.StatusPending
	ticket.TicketCheckStatus = models.CheckStatusNew

	if result := db.DB.Create(&ticket); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func GetTicket(c *gin.Context) {
	var tickets []models.Ticket

	if result := db.DB.Find(&tickets); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func UpdateTicketStatus(c *gin.Context) {
	id := c.Param("id")
	var ticket models.Ticket

	if result := db.DB.First(&ticket, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		return
	}

	var input struct {
		Status      models.TicketStatus      `json:"ticket_status"`
		CheckStatus models.TicketCheckStatus `json:"ticket_check_status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	// Update ticket
	ticket.TicketStatus = input.Status
	ticket.TicketCheckStatus = input.CheckStatus
	db.DB.Save(&ticket)
}

func DeleteTicket(c *gin.Context) {
	id := c.Param("id")
	var ticket models.Ticket

	if result := db.DB.First(&ticket, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
	db.DB.Delete(&ticket)
}
