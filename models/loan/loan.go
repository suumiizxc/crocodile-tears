package loan

import "time"

type Loan struct {
	ID             uint `json:"id" gorm:"primary_key"`
	ClientID       uint
	AccountID      uint
	RequestAmount  float32
	ApprovedAmount float32
	ProductID      uint
	MonthlyPayment float32
	DurationDay    uint
	Status         string
	CreateAt       time.Time
}
