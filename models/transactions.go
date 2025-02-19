package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transactions struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Description string    `json:"description" gorm:"size:50;not null"`
	Value       float64   `json:"value" gorm:"not null;check:value > 0"`
	Date        time.Time `json:"date" gorm:"not null"`
	Currency    string    `json:"currency" gorm:"size:3;not null"`
}

type TransactionsInput struct {
	Description string  `json:"description" binding:"required,max=50"`
	Value       float64 `json:"value" binding:"required,gt=0"`
	Date        string  `json:"date" binding:"required,datetime=2006-01-02"`
	Currency    string  `json:"currency" binding:"required,max=3"`
}

func (t *Transactions) BeforeCreate(d *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
