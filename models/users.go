package models

import (
	"gorm.io/gorm"
)

type Users struct {
	Rank       Ranks  `gorm:"embedded"`
	Name       string `json:"name"`
	EmployeeId int16  `json:"employeeid"`
	Password   string `json:"password"`
	gorm.Model
}
