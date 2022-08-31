package marketplace

import (
	"time"

	"github.com/suumiizxc/car-marketplace/config"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Name             string
	ManufactoryID    uint64 // Uildver
	MarkID           uint64 // mashinii torol prius
	LocationDetail   string // bairshil
	ConditionID      uint64 // Nohtsol dugaar avsan eseh
	TypeID           uint64 //torol jeep, suudliin, ger buliin
	DoorNumber       uint64
	SteeringWheel    uint64 // Khurd zov buruu
	WheelDriveID     uint64 // fwd, rwd, 4wd
	ManufacturedDate time.Time
	ImportedDate     time.Time
	EngineID         uint64  // hodolguur benzin, disel, hybrid
	MotorCapacity    float32 // motor bagtaamj 1.5, 1.8
	VelocityBoxID    uint64  // hurdnii hairtsag mechanic, automat
	InnerColorID     uint64
	ColorID          uint64
	WentDistance     float32
	LeasingTypeID    uint64
	Price            float32
	LocationID       uint64
	Images           []CarImage      `gorm:"foreignKey:CarID"`
	DiagonisImages   []DiagonisImage `gorm:"foreignKey:CarID"`
}

type CarImage struct {
	gorm.Model
	CarID uint64
	Image string
	Url   string
}

type DiagonisImage struct {
	gorm.Model
	CarID uint64
	Image string
	Url   string
}

func (c *Car) Migration() error {
	err := config.DB.AutoMigrate(
		Car{},
		CarCategory{},
		CarColor{},
		CarCondition{},
		CarEngine{},
		CarInnerColor{},
		CarLeasingType{},
		CarLocation{},
		CarManufactory{},
		CarMark{},
		CarType{},
		CarVelocityBox{},
		CarWheelDrive{},
		CarFeature{},
	)
	return err

}
