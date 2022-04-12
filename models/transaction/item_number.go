package transaction

type ItemNumber struct {
	ID          uint `json:"id" gorm:"primary_key"`
	Name        string
	Description string
}
