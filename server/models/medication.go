package models

import (
	"time"

	"gorm.io/gorm"
)

type Medication struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Serial      string `json:"serial"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
