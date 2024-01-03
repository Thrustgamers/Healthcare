package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDb() {

	db_user := "root"
	db_host := "127.0.0.1"
	db_port := "3306"
	db_name := "inventory"
	location_currentzone, location_err := time.LoadLocation("Local")

	if location_err != nil {
		log.Fatal(location_err)
	}

	connection := fmt.Sprintf("%s:@tcp(%s:%s)/%s?parseTime=true&loc=%s", db_user, db_host, db_port, db_name, location_currentzone)

	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to connect database")
	}

	//Migrate the schema
	// db.AutoMigrate(&models.Ranks{}, &models.Users{}, &models.Medication{})

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
