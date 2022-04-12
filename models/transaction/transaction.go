package transaction

import "time"

type Transaction struct {
	ID            uint `json:"id" gorm:"primary_key"`
	AccountID     uint
	LoanID        uint
	BankID        uint
	ItemNumberID  uint
	AccountNumber string
	Description   string
	Txndate       time.Time
	Amount        float32
	CreatedAt     time.Time
}
