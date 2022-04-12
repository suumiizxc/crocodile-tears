package account

type Currency struct {
	ID        uint `json:"id" gorm:"primary_key"`
	Name      string
	ShortName string
}
