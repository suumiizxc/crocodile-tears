package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/car-marketplace/config"
	client "github.com/suumiizxc/car-marketplace/controllers/client"
	"github.com/suumiizxc/car-marketplace/helper/redis"
)

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

	// Run the server
	r.Run()
}
