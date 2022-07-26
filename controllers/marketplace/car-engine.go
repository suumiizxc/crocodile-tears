package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarEngineInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCarEngine(c *gin.Context) {
	var input CreateCarEngineInput
	var engine marketplace.CarEngine

	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&engine, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := engine.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

func FindCarEngineByID(c *gin.Context) {
	var engine marketplace.CarEngine
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	engine.ID = id
	cm, err := engine.FindByID()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func GetCarEngineList(c *gin.Context) {
	var engine marketplace.CarEngine
	cms, err := engine.List()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cms})
}
