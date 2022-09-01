package marketplace

import (
	"fmt"

	"github.com/suumiizxc/car-marketplace/config"
)

// import "github.com/suumiizxc/car-marketplace/config"

type CarWheelDrive struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (cwd *CarWheelDrive) Create() error {
	return config.DB.Create(&cwd).Error
}

func (cwd *CarWheelDrive) FindByID() (CarWheelDrive, error) {
	cwdm := CarWheelDrive{}
	if err := config.DB.Find(&cwdm, cwd.ID).Error; err != nil {
		return CarWheelDrive{}, err
	}
	if cwdm.ID == 0 {
		return CarWheelDrive{}, fmt.Errorf("Not found")
	}
	return cwdm, nil
}

func (cwd *CarWheelDrive) List() ([]CarWheelDrive, error) {
	cwds := []CarWheelDrive{}
	if err := config.DB.Find(&cwds).Error; err != nil {
		return []CarWheelDrive{}, err
	}
	return cwds, nil
}

func (cwd *CarWheelDrive) DeleteByID() error {
	cwdm := CarWheelDrive{}
	if err := config.DB.Find(&cwdm, cwd.ID).Error; err != nil {
		return err
	}
	return nil
}
