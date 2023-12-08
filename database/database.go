package database

import (
	"api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDb() {

	// db_user := os.Getenv("DB_USER")
	// db_host := os.Getenv("DB_HOST")
	// db_port := os.Getenv("DB_PORT")
	// db_name := os.Getenv("DB_NAME")

	// connection := fmt.Sprintf("%s:@tcp(%s:%s)/%s", db_user, db_host, db_port, db_name)
	// connection := "root:@tcp(127.0.0.1:3306)/inventory"

	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/inventory?parseTime=true&loc=America%2FNew_York"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&models.Items{}, &models.Ranks{}, &models.Users{})

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

	Init()
}
