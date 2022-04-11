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
	r.DELETE("/core/customer/education-degree/:id/delete", customer.EducationDegreeDelete)

	r.POST("/core/customer/nation/create", customer.NationCreate)
	r.DELETE("/core/customer/nation/:id/delete", customer.NationDelete)
	r.GET("/core/customer/nation/:id", customer.NationGet)
	r.GET("/core/customer/nations/:limit/:page", customer.NationList)

	r.POST("/core/customer/ethnicity/create", customer.EthnicityCreate)
	r.DELETE("/core/customer/ethnicity/:id/delete", customer.EthnicityDelete)
	r.GET("/core/customer/ethnicity/:id", customer.EthnicityGet)
	r.GET("/core/customer/ethnicitys/:limit/:page", customer.EthnicityList)

	r.POST("/core/customer/profession/create", customer.ProfessionCreate)
	r.DELETE("/core/customer/profession/:id/delete", customer.ProfessionDelete)
	r.GET("/core/customer/profession/:id", customer.ProfessionGet)
	r.GET("/core/customer/professions/:limit/:page", customer.ProfessionList)

	r.POST("/core/customer/address_type/create", customer.AddressTypeCreate)
	r.DELETE("/core/customer/address_type/:id/delete", customer.AddressTypeDelete)
	r.GET("/core/customer/address_type/:id", customer.AddressTypeGet)
	r.GET("/core/customer/address_types/:limit/:page", customer.AddressTypeList)

	r.POST("/core/customer/address/create", customer.AddressCreate)
	r.DELETE("/core/customer/address/:id/delete", customer.AddressDelete)
	r.GET("/core/customer/address/:id", customer.AddressGet)
	r.GET("/core/customer/addresses/:limit/:page", customer.AddressList)

	r.POST("/core/customer/respect/create", customer.RespectCreate)
	r.DELETE("/core/customer/respect/:id/delete", customer.RespectDelete)
	r.GET("/core/customer/respect/:id", customer.RespectGet)
	r.GET("/core/customer/respects", customer.RespectList)

	// r.POST("/core/customer/soc/create", customer.SOCCreate)
	r.DELETE("/core/customer/soc/:id/delete", customer.SOCDelete)
	r.GET("/core/customer/soc/:id", customer.SOCGet)
	r.GET("/core/customer/socs", customer.SOCList)

	// r.POST("/core/customer/reltype/create", customer)
	r.DELETE("/core/customer/reltype/:id/delete", customer.RelTypeDelete)
	r.GET("/core/customer/reltype/:id", customer.RelTypeGet)
	r.GET("/core/customer/reltypes", customer.RelTypeList)

	//r.POST("/core/customer/customer_category/create", customer.CustomerCategoryCreate)
	r.DELETE("/core/customer/customer_category/:id/delete", customer.CustomerCategoryDelete)
	r.GET("/core/customer/customer_category/:id", customer.CustomerCategoryGet)
	r.GET("/core/customer/customer_categories", customer.CustomerCategoryList)

	//r.POST("/core/customer/crctype/create", customer.CustomerCategoryCreate)
	r.DELETE("/core/customer/crctype/:id/delete", customer.CRCTypeDelete)
	r.GET("/core/customer/crctype/:id", customer.CRCTypeGet)
	r.GET("/core/customer/crctypes", customer.CRCTypeList)

	// Run the server
	r.Run()
}
