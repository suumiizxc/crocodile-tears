package marketplace

import "github.com/suumiizxc/car-marketplace/config"

type CarLeasingType struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (cl *CarLeasingType) Create() error {
	return config.DB.Create(&cl).Error
}

func (cl *CarLeasingType) FindByID() (CarLeasingType, error) {
	var clm CarLeasingType
	if err := config.DB.Find(&clm, cl.ID).Error; err != nil {
		return CarLeasingType{}, err
	}
	return clm, nil
}

func (cl *CarLeasingType) List() ([]CarLeasingType, error) {
	var clms []CarLeasingType
	if err := config.DB.Find(&clms).Error; err != nil {
		return []CarLeasingType{}, err
	}
	return clms, nil
}
