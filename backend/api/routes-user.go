package api

import (
	"nabbr-appraisal/utils"

	"github.com/gin-gonic/gin"
)

func RoutesUser(router *gin.Engine) {
	user := router.Group("/api/v1/user")
	{
		// Set up authentication middleware
		user.Use(utils.Authentication())
		// All routes that require authentication should be placed after this line
		user.GET("/user", GetUsersAll())
		user.GET("/user/:userId", GetUser())
	}
}
