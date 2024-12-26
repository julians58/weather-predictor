package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/julians58/weather-predictor/internal/db"
	"github.com/julians58/weather-predictor/internal/models"
)

func GetAllWeather(c *gin.Context) {
	var weatherData []models.Weather
	result := db.DB.Find(&weatherData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, weatherData)
}

func GetWeather(c *gin.Context) {
	day, err := strconv.Atoi(c.Param("day"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid day"})
		return
	}

	var weather models.Weather
	result := db.DB.First(&weather, day)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Weather not found"})
		return
	}

	c.JSON(http.StatusOK, weather)
}

func GetWeatherStatistics(c *gin.Context) {
	var stats models.WeatherStatistics
	result := db.DB.Last(&stats)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Statistics not found"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func GetDaysWithCondition(c *gin.Context) {
	condition := c.Param("condition")

	var days []models.Weather
	result := db.DB.Where("condition = ?", condition).Find(&days)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, days)
}
