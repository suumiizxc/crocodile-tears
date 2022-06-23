package marketplace

import "gorm.io/gorm"

type CarCategory struct {
	gorm.Model
	Name        string
	HasChildren bool
}

type CarFeature struct {
	gorm.Model
	FeatureName  string
	FeatureType  string
	Required     bool
	MeasureUnit  string
	FeatureName2 string
}
