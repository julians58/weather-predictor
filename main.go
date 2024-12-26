package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/julians58/weather-predictor/internal/db"
	"github.com/julians58/weather-predictor/internal/routes"
)

func main() {
	log.Println("Starting server...")
	// Initialize the database
	if err := db.Init(); err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// Seed the database and wait until it's done
	if err := db.Seed(); err != nil {
		log.Fatalf("Could not seed database: %v", err)
	}

	// Set up the router
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
