package api

import (
	"github.com/gin-gonic/gin"
)

// UserRoutes function
func RoutesUser(router *gin.Engine) {
	user := router.Group("/api/v1/user")
	{
		user.POST("/users/signup", SignUp())
		user.POST("/users/login", Login())

	}
}
