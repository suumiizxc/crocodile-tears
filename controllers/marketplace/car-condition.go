package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarConditionInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCarCondition(c *gin.Context) {
	var input CreateCarConditionInput
	var condition marketplace.CarCondition
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&condition, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := condition.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "successfully created"})
}

func FindCarConditionByID(c *gin.Context) {
	var condition marketplace.CarCondition
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	condition.ID = id
	cm, err := condition.FindByID()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func GetCarConditionList(c *gin.Context) {
	var condition marketplace.CarCondition
	cms, err := condition.GetList()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cms})
}

func DeleteCarConditionByID(c *gin.Context) {
	var condition marketplace.CarCondition
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	condition.ID = id
	err := condition.DeleteByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": err})
}
