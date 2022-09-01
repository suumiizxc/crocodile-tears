package marketplace

import (
	"fmt"

	"github.com/suumiizxc/car-marketplace/config"
)

type CarCondition struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (cc *CarCondition) Create() error {
	return config.DB.Create(&cc).Error
}

func (cc *CarCondition) FindByID() (CarCondition, error) {
	var ccm CarCondition
	if err := config.DB.Find(&ccm, cc.ID).Error; err != nil {
		return CarCondition{}, err
	}
	if ccm.ID == 0 {
		return CarCondition{}, fmt.Errorf("Not found")
	}
	return ccm, nil
}

func (cc *CarCondition) GetList() ([]CarCondition, error) {
	var ccms []CarCondition
	if err := config.DB.Find(&ccms).Error; err != nil {
		return []CarCondition{}, err
	}
	return ccms, nil
}

func (cc *CarCondition) DeleteByID() error {
	var ccm CarCondition
	if err := config.DB.Delete(&ccm, cc.ID).Error; err != nil {
		return err
	}
	return nil
}
