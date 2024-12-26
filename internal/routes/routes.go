package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/julians58/weather-predictor/internal/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/weather/:day", handlers.GetWeather)
	router.GET("/days/:condition", handlers.GetDaysWithCondition)
	router.GET("/weather", handlers.GetAllWeather)
}
