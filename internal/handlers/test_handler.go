package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, GO KhanhRom"})
	}
}
