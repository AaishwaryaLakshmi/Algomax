package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupEnrollmentRoutes(router *gin.Engine) {
	enrollmentGroup := router.Group("/enrollments")
	{
		enrollmentGroup.POST("/", enrollCourse)
	}
}

func enrollCourse(c *gin.Context) {
	// Simplified enrollment logic for demonstration
	c.JSON(http.StatusOK, gin.H{"message": "Enrolled successfully"})
}
