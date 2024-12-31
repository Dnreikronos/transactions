package models

import (
	"time"

	"github.com/google/uuid"
)

type Transactions struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Description string    `json:"description"`
	Value       int64     `json:"value"`
	Date        time.Time `json:"date"`
}
