package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarManufactoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCarManufactory(c *gin.Context) {
	var input CreateCarManufactoryInput
	var manufactory marketplace.CarManufactory
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&manufactory, smapping.MapFields(input)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := manufactory.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

func FindCarManufactoryByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var manufactory marketplace.CarManufactory
	manufactory.ID = id
	cm, err := manufactory.FindByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func FindCarManufactoryByName(c *gin.Context) {
	name := c.Param("name")
	var manufactory marketplace.CarManufactory
	manufactory.Name = name
	cm, err := manufactory.FindByName()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func GetCarManufactoryList(c *gin.Context) {
	var manufactory marketplace.CarManufactory
	cms, err := manufactory.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cms})
}
