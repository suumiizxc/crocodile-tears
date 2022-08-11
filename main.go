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
	carFeatureRoute := r.Group("/car/feature")
	{
		carFeatureRoute.GET("/all", marketplace.FindCarFeatures)
		carFeatureRoute.GET("/get-by-id/:id", marketplace.FindCarFeatureById)
		carFeatureRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarFeature)
		carFeatureRoute.PUT("/update", custom_middleware.EnsureLoggedInAdmin(), marketplace.UpdateCarFeature)
	}
	carCategoryRoute := r.Group("/car/category")
	{
		carCategoryRoute.GET("/all", marketplace.FindCarCategories)
		carCategoryRoute.GET("/get-by-id/:id", marketplace.FindCarCategoryById)
		carCategoryRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarCategory)
		carCategoryRoute.PUT("/update", custom_middleware.EnsureLoggedInAdmin(), marketplace.UpdateCarCategory)
		carCategoryRoute.DELETE("/delete/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarCategory)
	}
	carColorRoute := r.Group("/car/color")
	{
		carColorRoute.GET("/all", marketplace.GetCarColorList)
		carColorRoute.GET("/get-by-id/:id", marketplace.FindCarColorByID)
		carColorRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarColor)
		carColorRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarColorByID)
	}
	carConditionRoute := r.Group("/car/condition")
	{
		carConditionRoute.GET("/all", marketplace.GetCarConditionList)
		carConditionRoute.GET("/get-by-id/:id", marketplace.FindCarConditionByID)
		carConditionRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarCondition)
		carConditionRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarConditionByID)
	}
	carEngineRoute := r.Group("/car/engine")
	{
		carEngineRoute.GET("/all", marketplace.GetCarEngineList)
		carEngineRoute.GET("/get-by-id/:id", marketplace.FindCarEngineByID)
		carEngineRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarEngine)
		carEngineRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarEngineByID)
	}
	carInnerColorRoute := r.Group("/car/inner-color")
	{
		carInnerColorRoute.GET("/all", marketplace.GetCarInnerColorList)
		carInnerColorRoute.GET("/get-by-id/:id", marketplace.FindCarInnerColorByID)
		carInnerColorRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarInnerColor)
		carInnerColorRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarInnerColorByID)
	}
	carLeasingTypeRoute := r.Group("/car/leasing-type")
	{
		carLeasingTypeRoute.GET("/all", marketplace.GetCarLeasingTypeList)
		carLeasingTypeRoute.GET("/get-by-id/:id", marketplace.FindCarLeasingTypeByID)
		carLeasingTypeRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarLeasingType)
		carLeasingTypeRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarLeasingTypeByID)
	}
	carLocationRoute := r.Group("/car/location")
	{
		carLocationRoute.GET("/all", marketplace.GetCarLocationList)
		carLocationRoute.GET("/get-by-id/:id", marketplace.FindCarLocationByID)
		carLocationRoute.GET("/get-by-name/:name", marketplace.FindCarLocationByName)
		carLocationRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarLocation)
		carLocationRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarLocationByID)
	}
	carManufactoryRoute := r.Group("/car/manufactory")
	{
		carManufactoryRoute.GET("/all", marketplace.GetCarManufactoryList)
		carManufactoryRoute.GET("/get-by-id/:id", marketplace.FindCarManufactoryByID)
		carManufactoryRoute.GET("/get-by-name/:name", marketplace.FindCarManufactoryByName)
		carManufactoryRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarManufactory)
		carManufactoryRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarManufactoryByID)
	}
	carMarkRoute := r.Group("/car/mark")
	{
		carMarkRoute.GET("/all", marketplace.GetCarMarkList)
		carMarkRoute.GET("/manufactory-by-id/:id", marketplace.FindCarMyMarkCMID)
		carMarkRoute.GET("/get-by-id/:id", marketplace.FindCarMarkByID)
		carMarkRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarMark)
		carMarkRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarMarkByID)
	}
	carTypeRoute := r.Group("/car/type")
	{
		carTypeRoute.GET("/all", marketplace.GetCarTypeList)
		carTypeRoute.GET("/get-by-id/:id", marketplace.FindCarTypeByID)
		carTypeRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarType)
		carTypeRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarTypeByID)
	}
	carVelocityRoute := r.Group("/car/velocity")
	{
		carVelocityRoute.GET("/all", marketplace.FindCarVelocityBoxList)
		carVelocityRoute.GET("/get-by-id/:id", marketplace.GetCarVelocityBoxByID)
		carVelocityRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarVelocityBox)
		carVelocityRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarVelocityBoxByID)
	}
	carWheelDriveRoute := r.Group("/car/wheel-drive")
	{
		carWheelDriveRoute.GET("/all", marketplace.GetCarWheelDriveList)
		carWheelDriveRoute.GET("/get-by-id/:id", marketplace.GetCarWheelDriveByID)
		carWheelDriveRoute.POST("/create", custom_middleware.EnsureLoggedInAdmin(), marketplace.CreateCarWheelDrive)
		carWheelDriveRoute.DELETE("/delete-by-id/:id", custom_middleware.EnsureLoggedInAdmin(), marketplace.DeleteCarWheelDriveByID)
	}
	// Run the server
	r.Run()
}
