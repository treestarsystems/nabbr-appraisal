package api

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// StartServer creates a new server instance
func StartServer() *gin.Engine {

	// This should not be hardcoded. It should be set in the environment but for some reason it is not working
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set up the router
	router := gin.Default()
	router.Use(gin.Logger())
	RoutesAuth(router)
	RoutesUser(router)
	RoutesAppraisal(router)

	router.Run()
	log.Printf("Server started on port :%s\n", os.Getenv("PORT"))

	return router
}
