package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transactions struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Description string    `json:"description"`
	Value       int64     `json:"value"`
	Date        time.Time `json:"date"`
}

type TransactionsInput struct {
	Description string `json:"description" binding:"required,gt=0"`
	Value       int64  `json:"value" binding:"required,gt=0"`
}

func (t *Transactions) BeforeCreate(d *gorm.DB) (err error) {
	t.ID = uuid.New()

	return
}
