package account

import "time"

type AccountHistory struct {
	ID           uint `json:"id" gorm:"primary_key"`
	AccountID    uint
	BegBal       float32
	EndBal       float32
	Txndate      time.Time
	Description  string
	ItemNumberID uint
	CreatedAt    time.Time
}
