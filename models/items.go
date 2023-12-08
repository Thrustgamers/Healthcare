package models

import (
	"gorm.io/gorm"
)

type Items struct {
	Name   string `json:"name"`
	Serial string `json:"serial"`
	gorm.Model
}
