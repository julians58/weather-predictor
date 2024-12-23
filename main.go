package main

import (
	"github.com/gin-gonic/gin"
	"github.com/julians58/weather-predictor/internal/db"
	"github.com/julians58/weather-predictor/internal/routes"
)

func main() {
	db.Init()
	db.Seed()
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
