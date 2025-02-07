package init

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// Define the database configuration
	dsn := "host=127.0.0.1 user=postgres password=doctorwho3210 dbname=golang_api_study port=5432"

	// Initialize the database connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}
