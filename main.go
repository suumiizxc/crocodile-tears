package main

import (
	"github.com/suumiizxc/gin-bookstore/config"
	"github.com/suumiizxc/gin-bookstore/controllers"
	client "github.com/suumiizxc/gin-bookstore/controllers/client"
	furniture "github.com/suumiizxc/gin-bookstore/controllers/furniture"
	"github.com/suumiizxc/gin-bookstore/helper/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	config.ConnectDatabase()
	redis.RedisConfig()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/furnitures", furniture.FindFurnitures)
	r.POST("/furnitures", furniture.CreateFurniture)

	r.GET("/clients", client.FindClients)
	r.POST("/client", client.CreateClient)
	r.POST("/client/login-phone", client.LoginPhone)
	r.POST("/client/login-email")
	r.GET("/client", client.ProfileClient)
	// Run the server
	r.Run()
}
