package main

import (
	"gohpapi/handlers"
)

func initializeRoutes() {

	// health
	router.GET("/health", handlers.GetHealth)

  // hit
	router.POST("/hit", handlers.PostHit)

}
