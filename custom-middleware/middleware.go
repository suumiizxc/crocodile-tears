package custom_middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/car-marketplace/helper/redis"
	models "github.com/suumiizxc/car-marketplace/models/client"
)

func EnsureLoggedInClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("access_token")
		objToken, err := redis.RS.Get(token).Result()
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not founded"})
			c.Abort()
			return
		}
		var client models.Client
		err = json.Unmarshal([]byte(objToken), &client)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not structured"})
			c.Abort()
			return
		}
		if client.Role == 1 {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func EnsureLoggedInAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("access_token")
		objToken, err := redis.RS.Get(token).Result()
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not founded"})
			c.Abort()
			return
		}
		var client models.Client
		err = json.Unmarshal([]byte(objToken), &client)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not structured"})
			c.Abort()
			return
		}
		if client.Role == 3 {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func EnsureLoggedInOperator() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("access_token")
		objToken, err := redis.RS.Get(token).Result()
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not founded"})
			c.Abort()
			return
		}
		var client models.Client
		err = json.Unmarshal([]byte(objToken), &client)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Token not structured"})
			c.Abort()
			return
		}
		if client.Role == 2 {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}
