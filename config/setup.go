package config

import (
	"fmt"
	"os"

	"github.com/suumiizxc/gin-bookstore/models"
	client "github.com/suumiizxc/gin-bookstore/models/client"
	furniture "github.com/suumiizxc/gin-bookstore/models/furniture"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf("Failed env : %v", err))
	}
	dsn := os.Getenv("DB")
	connstring := os.ExpandEnv(dsn)
	database, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(
		&models.Book{},
		&furniture.Furniture{},
		&client.Client{},
		&client.Permission{},
	)

	DB = database
}
