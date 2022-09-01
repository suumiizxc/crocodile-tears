package marketplace

import (
	"fmt"

	"github.com/suumiizxc/car-marketplace/config"
)

// import "github.com/suumiizxc/car-marketplace/config"

type CarMark struct {
	ID               uint64 `json:"id" gorm:"primary_key"`
	CarManufactoryID uint64 `json:"car_manufactory_id"`
	Name             string `json:"name"`
}

func (cm *CarMark) Create() error {
	err := config.DB.Create(&cm).Error
	return err
}

func (cm *CarMark) FindByID() (CarMark, error) {
	var cma CarMark
	if err := config.DB.Find(&cma, cm.ID).Error; err != nil {
		return CarMark{}, err
	}
	if cma.ID == 0 {
		return CarMark{}, fmt.Errorf("Not found")
	}
	return cma, nil
}

func (cm *CarMark) FindByCMID() ([]CarMark, error) {
	var cmas []CarMark
	if err := config.DB.Where("car_manufactory_id = ?", cm.CarManufactoryID).Find(&cmas).Error; err != nil {
		return []CarMark{}, err
	}
	return cmas, nil
}

func (cm *CarMark) GetList() ([]CarMark, error) {
	var cmas []CarMark
	if err := config.DB.Find(&cmas).Error; err != nil {
		return []CarMark{}, err
	}
	return cmas, nil
}

func (cm *CarMark) DeleteByID() error {
	var cma CarMark
	if err := config.DB.Delete(&cma, cm.ID).Error; err != nil {
		return err
	}
	return nil
}
