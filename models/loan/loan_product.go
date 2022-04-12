package loan

type LoanProduct struct {
	ID             uint `json:"id" gorm:"primary_key"`
	Name           string
	Description    string
	InterestRate   float32
	MinDurationDay uint
	MaxDurationDay uint
}
