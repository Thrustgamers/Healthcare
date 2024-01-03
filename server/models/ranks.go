package models

import (
	"time"

	"gorm.io/gorm"
)

type Ranks struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
