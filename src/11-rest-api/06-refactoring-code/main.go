package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()             // Initialize the database connection and create necessary tables
	server := gin.Default() // Create a new Gin router instance

	routes.RegisterRoutes(server) // Register the routes defined in the routes package

	server.Run(":8080") // localhost:8080
}
