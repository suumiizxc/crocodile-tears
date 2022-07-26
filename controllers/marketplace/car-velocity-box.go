package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarVelocityBoxInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCarVelocityBox(c *gin.Context) {
	var input CreateCarVelocityBoxInput
	var vebox marketplace.CarVelocityBox
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&vebox, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	if err := vebox.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

func GetCarVelocityBoxByID(c *gin.Context) {
	var vebox marketplace.CarVelocityBox
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	vebox.ID = id
	cm, err := vebox.GetByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func FindCarVelocityBoxList(c *gin.Context) {
	var veboxs marketplace.CarVelocityBox
	cms, err := veboxs.GetByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cms})
}
