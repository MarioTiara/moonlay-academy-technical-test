package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Mario Tiara",
		"bio":  "A Full Stack SOftware Enginner",
	})
}
