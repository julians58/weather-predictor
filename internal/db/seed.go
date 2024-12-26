package db

import (
	"fmt"
	"log"

	"github.com/julians58/weather-predictor/internal/models"
	"github.com/julians58/weather-predictor/internal/services"
)

// Seed seeds the database with weather data if it is empty.
func Seed() error {
	var count int64
	if err := DB.Model(&models.Weather{}).Count(&count).Error; err != nil {
		log.Printf("Error counting weather records: %v", err)
		return err
	}

	log.Printf("Current weather record count: %d", count)

	if count == 0 {
		log.Println("Seeding database with weather data...")
		maxPerimetro := 0.0
		rainBelowPeak := false
		var weatherData []models.Weather
		for dia := 0; dia <= 79; dia++ {
			weather, setPrevAsPeak, err := services.CalculateWeather(dia, &maxPerimetro, &rainBelowPeak)
			fmt.Println("macPerimetro", maxPerimetro)
			if err != nil {
				log.Printf("Error calculating weather for day %d: %v", dia, err)
				continue
			}
			if setPrevAsPeak {
				weatherData[len(weatherData)-1].Condition = "Pico de lluvia"
			}
			weatherData = append(weatherData, weather)

		}

		if len(weatherData) > 0 {
			if err := DB.Create(&weatherData).Error; err != nil {
				log.Printf("Error creating weather records: %v", err)
				return err
			}
			log.Printf("Successfully inserted %d weather records into the database", len(weatherData))
		} else {
			log.Println("No weather data to insert")
		}

		log.Println("Database seeded successfully!")
	} else {
		log.Println("Database already seeded!")
	}

	return nil
}
