package handlers

import (
	"gopkg.in/gin-gonic/gin.v1"
  "gohpapi/models"
	"net/http"
)

func PostHit(c *gin.Context) {

  var err error
	var json models.Message

	if err = c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "json decoding : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, json)
}
