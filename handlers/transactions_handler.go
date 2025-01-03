package handlers

import (
	"net/http"
	"time"

	"github.com/Dnreikronos/transactions/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTransaction(c *gin.Context) {
	var input models.TransactionsInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	transaction := models.Transactions{
		Description: input.Description,
		Value:       roundToCents(input.Value),
		Date:        parsedDate,
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func roundToCents(value float64) float64 {
	return float64(int(value*100+0.5)) / 100
}
