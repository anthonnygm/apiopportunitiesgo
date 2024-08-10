package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Config Routes
	configRoutes(router)

	// Run the server
	router.Run(":5000")
}