package loan

import "time"

type LoanHistory struct {
	ID            uint `json:"id" gorm:"primary_key"`
	LoanID        uint
	TransactionID uint
	Txndate       time.Time
	PrincPayment  float32
	IntPayment    float32
	BegBal        float32
	EndBal        float32
	Description   string
	Status        string
	CreatedAt     time.Time
}
