package marketplace

import (
	"net/http"

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
	if err := smapping.FillStruct(cwd, smapping.MapFields(input)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
}
