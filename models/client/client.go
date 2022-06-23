package client

import "time"

type Client struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Password     string `json:"password"`
	Registration string `json:"registration"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Token        string `gorm:"-" json:"token,omitempty"`
	IsActive     uint64 `json:"is_active"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
