package db

import (
	"encoding/json"
	"log"

	"github.com/julians58/weather-predictor/internal/models"
	"github.com/julians58/weather-predictor/internal/services"
)

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
		rainPeakDays := []int{}
		for dia := 0; dia <= 3650; dia++ {
			weather, setPrevAsPeak, err := services.CalculateWeather(dia, &maxPerimetro, &rainBelowPeak)
			if err != nil {
				log.Printf("Error calculating weather for day %d: %v", dia, err)
				continue
			}

			if setPrevAsPeak {
				weatherData[len(weatherData)-1].Condition = "Pico de lluvia"
				rainPeakDays = append(rainPeakDays, weatherData[len(weatherData)-1].Day)
			}

			weatherData = append(weatherData, weather)
		}

		if len(weatherData) > 0 {
			if err := DB.Create(&weatherData).Error; err != nil {
				log.Printf("Error creating weather records: %v", err)
				return err
			}
			log.Printf("Successfully inserted %d weather records into the database", len(weatherData))

			droughtPeriods := 0
			rainPeriods := 0
			optimalPeriods := 0
			isPrevNormal := false
			for i := 0; i < len(weatherData); i++ {
				if weatherData[i].Condition == "Sequia" {
					droughtPeriods++
				}

				if weatherData[i].Condition == "Lluvia" && weatherData[i-1].Condition == "Condiciones normales" {
					if weatherData[i].Condition == "Pico de lluvia" {
						rainPeakDays = append(rainPeakDays, weatherData[i].Day)
					}
					rainPeriods++
				}

				if weatherData[i].Condition == "Condiciones optimas de presion y temperatura" && isPrevNormal {
					optimalPeriods++
				}

				isPrevNormal = weatherData[i].Condition == "Condiciones normales"
			}

			rainPeakDaysJSON, err := json.Marshal(rainPeakDays)
			if err != nil {
				log.Printf("Error serializing rainPeakDays: %v", err)
				return err
			}

			stats := models.WeatherStatistics{
				DroughtPeriods: droughtPeriods,
				RainPeriods:    rainPeriods,
				OptimalPeriods: optimalPeriods,
				RainPeakDays:   string(rainPeakDaysJSON),
			}

			if err := DB.Create(&stats).Error; err != nil {
				log.Printf("Error creating weather statistics: %v", err)
				return err
			}
			log.Println("Successfully inserted weather statistics into the database")
		} else {
			log.Println("No weather data to insert")
		}

		log.Println("Database seeded successfully!")
	} else {
		log.Println("Database already seeded!")
	}

	return nil
}
