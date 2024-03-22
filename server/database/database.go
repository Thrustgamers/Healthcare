package database

import (
	"api/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

func getDB() (*gorm.DB, error) {
	// Preparing dotenv files
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err)
	}

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_driver := os.Getenv("DB_DRIVER")
	current_time := time.Now()
	location_currentzone, _ := current_time.Zone()

	if db_host == "" {
		return nil, errors.New("missing database information")
	}

	fmt.Printf("Database driver: %s found ", db_driver)

	var dsn string
	var driver gorm.Dialector
	switch db_driver {
	case "postgres":
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=%s", db_user, db_pass, db_host, db_port, db_name, location_currentzone)
		driver = postgres.Open(dsn)
	case "sqlServer":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", db_user, db_pass, db_host, db_port, db_name)
		driver = sqlserver.Open(dsn)
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_name)
		driver = mysql.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", db_driver)
	}

	// Open database connection
	db, err := gorm.Open(driver, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	return db, nil
}

var Database DbInstance

func ConnectToDb() {
	db, err := getDB()
	if err != nil {
		log.Error().Err(err)
		os.Exit(2)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Ranks{}, &models.Users{}, &models.Medication{})

	// Insert testing values
	db.FirstOrCreate(&models.Ranks{}, models.Ranks{Name: "Admin", Admin: "YES"})
	db.FirstOrCreate(&models.Users{}, models.Users{Rank: 1, Name: "Test", EmployeeId: 12345678, Password: "test"})

	Database = DbInstance{Db: db}
}
