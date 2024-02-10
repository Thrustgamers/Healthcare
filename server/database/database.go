package database

import (
	"api/models"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDb() {

	//Preparing dotenv files
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err)
	}

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	current_time := time.Now()
	location_currentzone, _ := current_time.Zone()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=%s", db_user, db_pass, db_host, db_port, db_name, location_currentzone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error().Err(err)
		os.Exit(2)
	}

	//Migrate the schema
	db.AutoMigrate(&models.Ranks{}, &models.Users{}, &models.Medication{})

	//Insert testing values
	db.FirstOrCreate(&models.Ranks{}, models.Ranks{Name: "Admin"})
	db.FirstOrCreate(&models.Users{}, models.Users{Rank: 1, Name: "Test", EmployeeId: 12345678, Password: "test"})

	Database = DbInstance{Db: db}
}
