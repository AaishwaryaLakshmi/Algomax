package routes

import (
	"github.com/gin-gonic/gin"
	"lms/models"
	"net/http"
)

func SetupCourseRoutes(router *gin.Engine) {
	courseGroup := router.Group("/courses")
	{
		courseGroup.GET("/", getCourses)
	}
}

func getCourses(c *gin.Context) {
	c.JSON(http.StatusOK, models.Courses)
}
