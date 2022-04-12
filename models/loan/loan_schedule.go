package loan

import "time"

type LoanSchedule struct {
	ID          uint `json:"id" gorm:"primary_key"`
	LoanID      uint
	Txndate     time.Time
	PrincAmount float32
	IntAmount   float32
}
