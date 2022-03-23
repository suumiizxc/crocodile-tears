package config

import (
	"os"

	"github.com/suumiizxc/gin-bookstore/models"
	furniture "github.com/suumiizxc/gin-bookstore/models/furniture"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "postgresql://root:root@0.0.0.0:5432/TA"
	connstring := os.ExpandEnv(dsn)
	database, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(
		&models.Book{},
		&furniture.Furniture{},
	)

	DB = database
}
