package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarLeasingTypeInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCarLeasingType(c *gin.Context) {
	var input CreateCarLeasingTypeInput
	var clt marketplace.CarLeasingType
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(clt, smapping.MapFields(input)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := clt.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

func FindCarLeasingTypeByID(c *gin.Context) {
	var leasingType marketplace.CarLeasingType
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	leasingType.ID = id
	cm, err := leasingType.FindByID()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func GetCarLeasingTypeList(c *gin.Context) {
	var leasingType marketplace.CarLeasingType
	cms, err := leasingType.List()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cms})
}
