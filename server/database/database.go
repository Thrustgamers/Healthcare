package database

import (
	"api/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDb() {

	// {
	// 	"label": "healthcare",
	// 	"host": "localhost",
	// 	"user": "postgres",
	// 	"port": 5432,
	// 	"ssl": false,
	// 	"database": "",
	// 	"password": "root"
	// }

	db_host := "localhost"
	db_user := "postgres"
	db_pass := "root"
	db_name := "healthcare"
	db_port := "5432"
	current_time := time.Now()
	location_currentzone, _ := current_time.Zone()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=%s", db_user, db_pass, db_host, db_port, db_name, location_currentzone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&models.Ranks{}, &models.Users{}, &models.Medication{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	// var product Product
	// db.First(&product, 1) // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(&product, 1)

	Database = DbInstance{Db: db}

	// Init()
}
