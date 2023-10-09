package main

import (
	"blogging-platform/database"
	"blogging-platform/handlers"
	"blogging-platform/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware
	router := gin.Default()

	// Initialize the database connection
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Failed to initialize the database:", err)
		return
	}

	// Defer the disconnection of the database client to ensure it's done when the application exits
	defer database.Disconnect()

	// Set the database instance in the handlers package
	handlers.SetDB(db)

	// Setup API routes
	routes.SetupPostRoutes(&router.RouterGroup)
	routes.SetupPostsRoutes(&router.RouterGroup)

	// Start the application and listen on port 8080
	router.Run(":8080")
}
