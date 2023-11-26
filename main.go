package main

import (
	"github.com/gin-gonic/gin"
	"lms/routes"
)

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Set up routes
	routes.SetupUserRoutes(router)
	routes.SetupCourseRoutes(router)
	routes.SetupEnrollmentRoutes(router)

	// Start the server
	router.Run(":8080")
}
