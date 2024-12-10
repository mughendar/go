package main

import (
	"log"
	"my-gin-mysql-app/config"
	"my-gin-mysql-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// Initialize the router
	router := gin.Default()

	// Setup routes
	routes.SetupUserRoutes(router, db)

	// Start the server
	router.Run(":8080")
}
