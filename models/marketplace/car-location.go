package marketplace

import "github.com/suumiizxc/car-marketplace/config"

type CarLocation struct {
	ID        uint64  `json:"id" gorm:"primary_key"`
	Name      string  `json:"name"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

func (cl *CarLocation) Create() error {
	return config.DB.Create(&cl).Error
}

func (cl *CarLocation) FindByID() (CarLocation, error) {
	var clm CarLocation
	if err := config.DB.Find(&clm, cl.ID).Error; err != nil {
		return CarLocation{}, err
	}
	return clm, nil
}

func (cl *CarLocation) FindByName() (CarLocation, error) {
	var clm CarLocation
	if err := config.DB.Where("name = ?", cl.Name).Find(&clm).Error; err != nil {
		return CarLocation{}, err
	}
	return clm, nil
}

func (cl *CarLocation) List() ([]CarLocation, error) {
	var clms []CarLocation
	if err := config.DB.Find(&clms).Error; err != nil {
		return []CarLocation{}, err
	}
	return clms, nil
}
