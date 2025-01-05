package handlers

import (
	"fmt"
	"net/http"

	"github.com/Dnreikronos/transactions/models"
	"github.com/Dnreikronos/transactions/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTransactionInCurrency(c *gin.Context) {
	id := c.Param("id")
	currency := c.Query("currency")

	if currency == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Currency query parameter is required"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var transaction models.Transactions
	if err := db.First(&transaction, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	fmt.Printf("Fetching exchange rate for transaction ID %s with currency %s on date %s\n", transaction.ID, currency, transaction.Date.Format("2006-01-02"))

	exchangeRate, err := utils.FetchExchangeRates(currency, transaction.Date)
	if err != nil {

		fmt.Printf("Failed to fetch exchange rate: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch exchange rate: %v", err)})
		return
	}

	convertedValue := roundToCents(transaction.Value * exchangeRate)

	c.JSON(http.StatusOK, gin.H{
		"transaction_id":     transaction.ID,
		"description":        transaction.Description,
		"date":               transaction.Date,
		"original_value_usd": transaction.Value,
		"exchange_rate":      exchangeRate,
		"converted_value":    convertedValue,
		"currency":           currency,
	})
}
