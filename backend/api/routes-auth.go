package api

import (
	"github.com/gin-gonic/gin"
)

// AuthRoutes function
func RoutesAuth(router *gin.Engine) {
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/signup", SignUp)
		auth.POST("/login", Login)
	}
}
