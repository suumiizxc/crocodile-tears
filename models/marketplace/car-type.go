package marketplace

import "github.com/suumiizxc/car-marketplace/config"

type CarType struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (ct *CarType) Create() error {
	return config.DB.Create(&ct).Error
}

func (ct *CarType) FindByID() (CarType, error) {
	var ctm CarType
	if err := config.DB.Find(&ctm, ct.ID).Error; err != nil {
		return CarType{}, err
	}
	return ctm, nil
}

func (ct *CarType) List() ([]CarType, error) {
	var cmtss []CarType
	if err := config.DB.Find(&cmtss).Error; err != nil {
		return []CarType{}, err
	}
	return cmtss, nil
}
