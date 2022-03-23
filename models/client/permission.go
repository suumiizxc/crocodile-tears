package client

type Permission struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	ClientID uint   `json:"client_id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Status   string `json:"status"`
}
