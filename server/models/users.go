package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID         uint   `gorm:"primaryKey"`
	Rank       Ranks  `gorm:"embedded"`
	Name       string `json:"name"`
	EmployeeId int    `json:"employeeid"`
	Password   string `json:"password"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
