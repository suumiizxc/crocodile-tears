package models

import "time"

type Furniture struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Measurement string  `json:"measurement"`
	Price       float32 `json:"price"`
	CategoryID  uint    `json:"category_id"`
	Status      string  `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
