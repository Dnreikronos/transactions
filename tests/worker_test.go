package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Dnreikronos/transactions/models"
	"github.com/Dnreikronos/transactions/queue"
	"github.com/Dnreikronos/transactions/worker"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestWorkerProcessesTransaction(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err, "Expected database connection to succeed")

	db.AutoMigrate(&models.Transactions{})

	go worker.StartTransactionWorker(db)

	transaction := models.Transactions{
		Description: "Test transaction",
		Value:       100.00,
		Date:        time.Now(),
	}
	transactionJSON, err := json.Marshal(transaction)
	assert.NoError(t, err, "Expected JSON marshaling to succeed")

	err = queue.PublishToQueue(transactionJSON)
	assert.NoError(t, err, "Expected transaction to be published to the queue")

	time.Sleep(1 * time.Second)

	var savedTransaction models.Transactions
	err = db.First(&savedTransaction).Error

	assert.NoError(t, err, "Expected transaction to be saved in the database")
	assert.Equal(t, transaction.Description, savedTransaction.Description)
	assert.Equal(t, transaction.Value, savedTransaction.Value)
	assert.True(t, transaction.Date.Equal(savedTransaction.Date))
}
