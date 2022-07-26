package marketplace

import "github.com/suumiizxc/car-marketplace/config"

type CarEngine struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (ce *CarEngine) Create() error {
	return config.DB.Create(&ce).Error
}

func (ce *CarEngine) FindByID() (CarEngine, error) {
	var cem CarEngine
	if err := config.DB.Find(&cem, ce.ID).Error; err != nil {
		return CarEngine{}, err
	}
	return cem, nil
}

func (ce *CarEngine) List() ([]CarEngine, error) {
	var cems []CarEngine
	if err := config.DB.Find(&cems).Error; err != nil {
		return []CarEngine{}, err
	}
	return cems, nil
}
