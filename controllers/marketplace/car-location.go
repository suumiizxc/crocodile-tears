package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarLocationInput struct {
	Name      string  `json:"name" binding:"required"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

func CreateCarLocation(c *gin.Context) {
	var input CreateCarLocationInput
	var location marketplace.CarLocation
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&location, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := location.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

func FindCarLocationByID(c *gin.Context) {
	var location marketplace.CarLocation
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	location.ID = id
	cm, err := location.FindByID()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func FindCarLocationByName(c *gin.Context) {
	var location marketplace.CarLocation
	name := c.Param("name")
	location.Name = name
	cm, err := location.FindByName()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func GetCarLocationList(c *gin.Context) {
	var location marketplace.CarLocation
	cms, err := location.List()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cms})
}
