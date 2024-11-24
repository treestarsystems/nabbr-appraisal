package api

import (
	"fmt"
	"nabbr-appraisal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authz validates token and authorizes users
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    "error - Authentication: No Authorization header provided",
				"payload":    []string{},
			})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":     "failure",
				"httpStatus": http.StatusInternalServerError,
				"message":    fmt.Sprintf("error - Authentication: No Authorization header provided %v", err),
				"payload":    []string{},
			})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("firstName", claims.FirstName)
		c.Set("lastName", claims.LastName)
		c.Set("userId", claims.UserId)
		c.Next()
	}
}
