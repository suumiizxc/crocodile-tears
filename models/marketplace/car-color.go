package marketplace

import "github.com/suumiizxc/car-marketplace/config"

type CarColor struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (cc *CarColor) Create() error {
	return config.DB.Create(&cc).Error
}

func (cc *CarColor) GetByID() (CarColor, error) {
	var ccm CarColor
	if err := config.DB.Find(&ccm, cc.ID).Error; err != nil {
		return CarColor{}, err
	}
	return ccm, nil
}

func (cc *CarColor) List() ([]CarColor, error) {
	var ccm []CarColor
	if err := config.DB.Find(&ccm).Error; err != nil {
		return []CarColor{}, err
	}
	return ccm, nil
}
