package marketplace

import (
	"fmt"

	"github.com/suumiizxc/car-marketplace/config"
)

// import "github.com/suumiizxc/car-marketplace/config"

type CarVelocityBox struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (cvb *CarVelocityBox) Create() error {
	return config.DB.Create(&cvb).Error
}

func (cvb *CarVelocityBox) GetByID() (CarVelocityBox, error) {
	var cvbm CarVelocityBox
	if err := config.DB.Find(&cvbm, cvb.ID).Error; err != nil {
		return CarVelocityBox{}, err
	}
	if cvbm.ID == 0 {
		return CarVelocityBox{}, fmt.Errorf("Not found")
	}
	return cvbm, nil
}

func (cvb *CarVelocityBox) List() ([]CarVelocityBox, error) {
	var cvbms []CarVelocityBox
	if err := config.DB.Find(&cvbms).Error; err != nil {
		return []CarVelocityBox{}, err
	}
	return cvbms, nil
}

func (cvb *CarVelocityBox) DeleteByID() error {
	var cvbm CarVelocityBox
	if err := config.DB.Find(&cvbm, cvb.ID).Error; err != nil {
		return err
	}
	return nil
}
