package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID         uint `gorm:"primaryKey"`
	Rank       int
	RankRefer  Ranks  `gorm:"foreignKey:Rank"`
	Name       string `json:"name"`
	EmployeeId int    `json:"employeeid"`
	Password   string `json:"password"`
	Admin      string `gorm:"default:'NO'" json:"admin"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
