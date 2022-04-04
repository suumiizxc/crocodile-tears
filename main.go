package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suumiizxc/gin-bookstore/config"
	"github.com/suumiizxc/gin-bookstore/controllers"
	client "github.com/suumiizxc/gin-bookstore/controllers/client"
	customer "github.com/suumiizxc/gin-bookstore/controllers/core/customer"
	furniture "github.com/suumiizxc/gin-bookstore/controllers/furniture"
	helper_core "github.com/suumiizxc/gin-bookstore/helper/core"
	"github.com/suumiizxc/gin-bookstore/helper/redis"
)

func main() {
	r := gin.Default()

	// Connect to database

	config.ConnectDatabase()
	helper_core.CH.Init()
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
	r.POST("/client/login-email", client.LoginEmail)
	r.GET("/client", client.ProfileClient)

	r.POST("/core/customer/create-test", customer.CreateTest)
	r.POST("/core/customer/create", customer.CreateCustomer)
	r.POST("/core/customer/countryCodes/:limit/:page", customer.GetCountryCodes)

	r.GET("/core/customer/education-degrees", customer.EducationDegreeList)
	r.GET("/core/customer/education-degree/:id", customer.EducationDegreeGet)
	r.POST("/core/customer/education-degree/create", customer.EducationDegreeCreate)
	r.DELETE("/core/customer/education-degree/delete/:id", customer.EducationDegreeDelete)

	r.GET("/core/customer/education-levels", customer.EducationLevelList)
	r.GET("/core/customer/education-level/:id", customer.EducationLevelGet)
	r.POST("/core/customer/education-level/create", customer.EducationLevelCreate)
	r.DELETE("/core/customer/education-level/delete/:id", customer.EducationLevelDelete)

	// Run the server
	r.Run()
}
