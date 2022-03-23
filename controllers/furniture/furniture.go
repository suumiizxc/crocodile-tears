package furniture

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/gin-bookstore/config"
	models "github.com/suumiizxc/gin-bookstore/models/furniture"
)

type CreateFurnitureInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Measurement string  `json:"measurement"`
	Price       float32 `json:"price"`
	CategoryID  uint    `json:"category_id"`
	Status      string  `json:"status"`
}

type UpdateFurnitureInput struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Measurement string  `json:"measurement"`
	Price       float32 `json:"price"`
	CategoryID  uint    `json:"category_id"`
	Status      string  `json:"status"`
}

func FindFurnitures(c *gin.Context) {
	var furnitures []models.Furniture
	config.DB.Find(&furnitures)
	c.JSON(http.StatusOK, gin.H{"data": furnitures})

}

func FindFurniture(c *gin.Context) {
	var furniture models.Furniture
	if err := config.DB.Where("id = ?", c.Param("id")).First(&furniture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": furniture, "message": "Successfully"})
}

func CreateFurniture(c *gin.Context) {
	var input CreateFurnitureInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	furniture := models.Furniture{}
	errDTO := smapping.FillStruct(&furniture, smapping.MapFields(&input))
	if errDTO != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errDTO})
	}
	if err := config.DB.Create(&furniture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": furniture})
}

func UpdateFurniture(c *gin.Context) {
	var input UpdateFurnitureInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := config.DB.Model(&models.Furniture{}).Updates(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteFurniture(c *gin.Context) {
	var furniture models.Furniture
	if err := config.DB.Where("id = ?", c.Param("id")).First(&furniture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	config.DB.Delete(&furniture)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
