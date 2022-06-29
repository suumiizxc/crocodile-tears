package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/config"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarCategoryInput struct {
	Name        string `json:"name" binding:"required"`
	HasChildren bool   `json:"has_children" binding:"required"`
}

type UpdateCarCategoryInput struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	HasChildren bool   `json:"has_children" binding:"required"`
}

type CreateCarFeatureInput struct {
	FeatureName  string `json:"feature_name" binding:"required"`
	FeatureType  string `json:"feature_type" binding:"required"`
	Required     bool   `json:"required" binding:"required"`
	MeasureUnit  string `json:"measure_unit"`
	FeatureName2 string `json:"feature_name2"`
}

type UpdateCarFeatureInput struct {
	ID           uint   `json:"id"`
	FeatureName  string `json:"feature_name" binding:"required"`
	FeatureType  string `json:"feature_type" binding:"required"`
	Required     bool   `json:"required" binding:"required"`
	MeasureUnit  string `json:"measure_unit"`
	FeatureName2 string `json:"feature_name2"`
}

func FindCarFeatures(c *gin.Context) {
	var features marketplace.CarFeature
	if err := config.DB.Find(&features).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": features})
}

func FindCarFeatureById(c *gin.Context) {
	var feature marketplace.CarFeature
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := config.DB.Find(&feature, id).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": feature})
}

func CreateCarFeature(c *gin.Context) {
	var input CreateCarFeatureInput
	var feature marketplace.CarFeature
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&feature, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&feature).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": feature})
}

func UpdateCarFeature(c *gin.Context) {
	var input UpdateCarFeatureInput
	var feature marketplace.CarFeature

	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&feature, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Updates(&feature).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": feature})
}

func FindCarCategories(c *gin.Context) {
	var categories []marketplace.CarCategory
	if err := config.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func FindCarCategoryById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var category marketplace.CarCategory
	if err := config.DB.Find(&category, id).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func CreateCarCategory(c *gin.Context) {
	var input CreateCarCategoryInput
	var category marketplace.CarCategory

	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}
	if err := smapping.FillStruct(&category, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": category})

}
