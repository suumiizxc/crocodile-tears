package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarMarkInput struct {
	CarManufactoryID uint64 `json:"car_manufactory_id" binding:"required"`
	Name             string `json:"name" binding:"required"`
}

func CreateCarMark(c *gin.Context) {
	var input CreateCarMarkInput
	var mark marketplace.CarMark
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&mark, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	if err := mark.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

func FindCarMarkByID(c *gin.Context) {
	var mark marketplace.CarMark
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	mark.ID = id
	cm, err := mark.FindByID()
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func FindCarMyMarkCMID(c *gin.Context) {
	var mark marketplace.CarMark
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	mark.CarManufactoryID = id
	cm, err := mark.FindByCMID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func GetCarMarkList(c *gin.Context) {
	var marks marketplace.CarMark
	cm, err := marks.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})

}

func DeleteCarMarkByID(c *gin.Context) {
	var mark marketplace.CarMark
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	mark.ID = id
	err := mark.DeleteByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}
