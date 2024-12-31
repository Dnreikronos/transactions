package migrations

import (
	m "github.com/Dnreikronos/transactions/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	createTables(db)
}

func createTables(db *gorm.DB) {
	db.AutoMigrate(&m.Transactions{})
}
