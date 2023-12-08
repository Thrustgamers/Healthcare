package database

import (
	"api/models"
	"log"
)

func Init() {

	item := models.Items{Name: "test", Serial: "512135435"}

	db := Database.Db.Where(&models.Items{Serial: "512135435"}).FirstOrCreate(&item)

	if db.Error != nil {
		log.Fatal(db.Error)
	}
}
