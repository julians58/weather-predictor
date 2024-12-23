package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/julians58/weather-predictor/internal/db"
	"github.com/julians58/weather-predictor/internal/models"
)

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
