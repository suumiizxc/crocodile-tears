package marketplace

import "github.com/suumiizxc/car-marketplace/config"

// import "github.com/suumiizxc/car-marketplace/config"

type CarInnerColor struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (cic *CarInnerColor) Create() error {
	return config.DB.Create(&cic).Error
}

func (cic *CarInnerColor) GetByID() (CarInnerColor, error) {
	var cicm CarInnerColor
	if err := config.DB.Find(&cicm, cic.ID).Error; err != nil {
		return CarInnerColor{}, err
	}
	return cicm, nil
}

func (cic *CarInnerColor) List() ([]CarInnerColor, error) {
	var cics []CarInnerColor
	if err := config.DB.Find(&cics).Error; err != nil {
		return []CarInnerColor{}, err
	}
	return cics, nil
}

func (cic *CarInnerColor) DeleteByID() error {
	var cicm CarInnerColor
	if err := config.DB.Delete(&cicm, cic.ID).Error; err != nil {
		return err
	}
	return nil
}
