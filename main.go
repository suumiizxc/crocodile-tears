package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/car-marketplace/config"
	client "github.com/suumiizxc/car-marketplace/controllers/client"
	marketplace "github.com/suumiizxc/car-marketplace/controllers/marketplace"
	custom_middleware "github.com/suumiizxc/car-marketplace/custom-middleware"

	"github.com/suumiizxc/car-marketplace/helper/redis"
)

// func ensureLoggedIn() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.GetHeader("access_token")
// 		if token == "2" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "pisda"})
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

func main() {
	r := gin.Default()

	// Connect to database

	config.ConnectDatabase()
	redis.RedisConfig()

	// Routes
	r.GET("/clients", client.FindClients)
	r.POST("/client", client.CreateClient)
	r.POST("/client/login-phone", client.LoginPhone)
	r.POST("/client/login-email", client.LoginEmail)

	r.GET("/client", client.ProfileClient)
	clientRoute := r.Group("/client")
	{
		clientRoute.GET("/profile", custom_middleware.EnsureLoggedInClient(), client.ProfileClient)
	}
	carFeatureRoute := r.Group("/car-feature")
	{
		carFeatureRoute.GET("/all", marketplace.FindCarFeatures)
		carFeatureRoute.GET("/get-by-id/:id", marketplace.FindCarFeatureById)
		carFeatureRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarFeature)
		carFeatureRoute.PUT("/update", custom_middleware.EnsureLoggedInAdmin(), marketplace.UpdateCarFeature)
	}
	// Run the server
	r.Run()
}
