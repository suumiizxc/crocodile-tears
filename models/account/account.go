package account

import "time"

type Account struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	CustCode   string `json:"custCode"`
	ClientID   uint
	Name       string
	Owner      string
	CurrencyID uint
	CreatedAt  time.Time
	ExpiredAt  time.Time
}
