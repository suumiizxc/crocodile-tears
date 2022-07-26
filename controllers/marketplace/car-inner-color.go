package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarInnerColorInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCarInnerColor(c *gin.Context) {
	var input CreateCarInnerColorInput
	var innerColor marketplace.CarInnerColor
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&innerColor, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := innerColor.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

func FindCarInnerColorByID(c *gin.Context) {
	var innerColor marketplace.CarInnerColor
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	innerColor.ID = id
	cm, err := innerColor.GetByID()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func GetCarInnerColorList(c *gin.Context) {
	var innerColor marketplace.CarInnerColor
	cms, err := innerColor.List()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cms})
}
