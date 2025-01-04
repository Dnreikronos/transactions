package worker

import (
	"encoding/json"
	"log"

	"github.com/Dnreikronos/transactions/models"
	"github.com/Dnreikronos/transactions/queue"
	"gorm.io/gorm"
)

func StartTransactionWorker(db *gorm.DB) {
	process := func(message []byte) {
		var transaction models.Transactions
		if err := json.Unmarshal(message, &transaction); err != nil {
			log.Printf("Failed to parse message: %v", err)
			return
		}

		if err := db.Create(&transaction).Error; err != nil {
			log.Printf("Failed to persist transaction: %v", err)
		} else {
			log.Printf("Transaction processed: %+v", transaction)
		}
	}

	queue.ConsumeQueue(process)
}
