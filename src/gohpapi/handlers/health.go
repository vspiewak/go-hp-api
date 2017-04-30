package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{"status": "UP"})
}
