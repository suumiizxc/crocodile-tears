package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarWheelDriveInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCarWheelDrive(c *gin.Context) {
	var input CreateCarWheelDriveInput
	var cwd marketplace.CarWheelDrive
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&cwd, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	if err := cwd.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully created"})
}

func GetCarWheelDriveByID(c *gin.Context) {
	var cwd marketplace.CarWheelDrive
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	cwd.ID = id
	cm, err := cwd.FindByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cm})
}

func GetCarWheelDriveList(c *gin.Context) {
	var cwd marketplace.CarWheelDrive
	cms, err := cwd.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cms})
}

func DeleteCarWheelDriveByID(c *gin.Context) {
	var cwd marketplace.CarWheelDrive
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	cwd.ID = id
	err := cwd.DeleteByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}
