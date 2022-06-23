package marketplace

import (
	"github.com/gin-gonic/gin"
	// client "github.com/suumiizxc/car-marketplace/controllers/client"
)

// var (
// 	module_name     = "marketplace"
// 	sub_module_name = "car_category"
// )

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

}
