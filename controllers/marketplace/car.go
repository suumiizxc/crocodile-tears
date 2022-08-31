package marketplace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/suumiizxc/car-marketplace/config"
	"github.com/suumiizxc/car-marketplace/models/marketplace"
)

type CreateCarInput struct {
	Name           string  `json:"name" binding:"required"`
	ManufactoryID  uint64  `json:"manufactory_id" binding:"required"`
	MarkID         uint64  `json:"mark_id" binding:"required"`
	LocationDetail string  `json:"location_detail" binding:"required"`
	ConditionID    uint64  `json:"condition_id" binding:"required"`
	TypeID         uint64  `json:"type_id" binding:"required"`
	DoorNumber     uint64  `json:"door_number" binding:"required"`
	SteeringWheel  uint64  `json:"steering_wheel" binding:"required"`
	WheelDriveID   uint64  `json:"wheel_drive_id" binding:"required"`
	EngineID       uint64  `json:"engine_id" binding:"required"`
	MotorCapacity  float32 `json:"motor_capacity" binding:"required"`
	VelocityBoxID  uint64  `json:"velocity_box_id" binding:"required"`
	InnerColorID   uint64  `json:"inner_color_id" binding:"required"`
	ColorID        uint64  `json:"color_id" binding:"required"`
	WentDistance   float32 `json:"went_distance" binding:"required"`
	LeasingTypeID  uint64  `json:"leasing_type_id" binding:"required"`
	Price          float32 `json:"price" binding:"required"`
	LocationID     uint64  `json:"location_id" binding:"required"`
}

type UpdateCarInput struct {
	ID             uint64  `json:"id" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	ManufactoryID  uint64  `json:"manufactory_id" binding:"required"`
	MarkID         uint64  `json:"mark_id" binding:"required"`
	LocationDetail string  `json:"location_detail" binding:"required"`
	ConditionID    uint64  `json:"condition_id" binding:"required"`
	TypeID         uint64  `json:"type_id" binding:"r:equired"`
	DoorNumber     uint64  `json:"door_number" binding:"required"`
	SteeringWheel  uint64  `json:"steering_wheel" binding:"required"`
	WheelDriveID   uint64  `json:"wheel_drive_id" binding:"required"`
	EngineID       uint64  `json:"engine_id" binding:"required"`
	MotorCapacity  float32 `json:"motor_capacity" binding:"required"`
	VelocityBoxID  uint64  `json:"velocity_box_id" binding:"required"`
	InnerColorID   uint64  `json:"inner_color_id" binding:"required"`
	ColorID        uint64  `json:"color_id" binding:"required"`
	WentDistance   float32 `json:"went_distance" binding:"required"`
	LeasingTypeID  uint64  `json:"leasing_type_id" binding:"required"`
	Price          float32 `json:"price" binding:"required"`
	LocationID     uint64  `json:"location_id" binding:"required"`
}

func FindCars(c *gin.Context) {
	var cars []marketplace.Car
	if err := config.DB.Find(&cars).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cars})
}

func FindCarById(c *gin.Context) {
	var car marketplace.Car
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := config.DB.Find(&car, id).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})
}

func CreateCar(c *gin.Context) {
	var input CreateCarInput
	var car marketplace.Car
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}

	if err := smapping.FillStruct(&car, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&car).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})
}

func UpdateCar(c *gin.Context) {
	var input UpdateCarInput
	var car marketplace.Car
	if errDTO := c.ShouldBind(&input); errDTO != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": errDTO.Error()})
		return
	}

	if err := smapping.FillStruct(&car, smapping.MapFields(&input)); err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Updates(&car).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})
}

func CarMigration(c *gin.Context) {
	car := marketplace.Car{}
	err := car.Migration()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully car migrated"})
}
