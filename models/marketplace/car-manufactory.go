package marketplace

import "github.com/suumiizxc/car-marketplace/config"

type CarManufactory struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (cm *CarManufactory) Create() error {
	return config.DB.Create(&cm).Error
}

func (cm *CarManufactory) FindByID() (CarManufactory, error) {
	var manu CarManufactory
	if err := config.DB.Find(&manu, cm.ID).Error; err != nil {
		return CarManufactory{}, err
	}
	return manu, nil
}

func (cm *CarManufactory) FindByName() (CarManufactory, error) {
	var manu CarManufactory
	if err := config.DB.Where("name = ?", cm.Name).Find(&manu).Error; err != nil {
		return CarManufactory{}, err
	}
	return manu, nil
}

func (cm *CarManufactory) List() ([]CarManufactory, error) {
	var manus []CarManufactory
	if err := config.DB.Find(&manus).Error; err != nil {
		return []CarManufactory{}, err
	}
	return manus, nil
}
