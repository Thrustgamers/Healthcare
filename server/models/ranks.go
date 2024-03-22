package models

import (
	"time"

	"gorm.io/gorm"
)

type Ranks struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Admin     string `gorm:"default: 'NO' json:"admin"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
