package db

import (
	"github.com/julians58/weather-predictor/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection and migrates the models.
func Init() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("weather.db"), &gorm.Config{})
	if err != nil {
		return err // Return the error to the caller
	}

	// Auto-migrate the Weather model
	if err := DB.AutoMigrate(&models.Weather{}); err != nil {
		return err // Return the error if migration fails
	}

	return nil // Successfully initialized and migrated
}
