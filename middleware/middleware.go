package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("access_token")
		if token != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "pisda"})
			return
		}
		c.Next()
	}
}
