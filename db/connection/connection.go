package connection

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenConnection() {
	var err error
	DB, err = gorm.Open(sqlite.Open("transactions.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected sucessfully!")
}
