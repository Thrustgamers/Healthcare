package models

import (
	"gorm.io/gorm"
)

type Ranks struct {
	Name string `json:"name"`
	gorm.Model
}
