package routes

import (
	"github.com/gin-gonic/gin"
	"lms/models"
	"net/http"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/login", loginUser)
	}
}

func loginUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Validate credentials (simplified for demonstration)
	for _, u := range models.Users {
		if u.Username == user.Username && u.Password == user.Password {
			c.JSON(http.StatusOK, gin.H{"role": u.Role})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
