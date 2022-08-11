package marketplace

import (
	"testing"
)

type CarColorTest struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func TestCreate(t *testing.T) {
	var carColor CarColor
	carColor.Name = "test-name"
	err := carColor.Create()
	// fmt.Println("error : ", err.Error())
	if err != nil {
		t.Error("Result failed be create car color, ", err.Error())

	}

}
