package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

var router *gin.Engine

func main() {

	// Set the default gin router
	router = gin.Default()

	// Initialize the routes
	initializeRoutes()

	// Serving application
	router.Run(":8080")

}
