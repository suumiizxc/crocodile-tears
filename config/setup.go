package config

import (
	"os"

	"github.com/suumiizxc/gin-bookstore/models"
	client "github.com/suumiizxc/gin-bookstore/models/client"
	furniture "github.com/suumiizxc/gin-bookstore/models/furniture"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "postgresql://suumiizxc:JS4Fr38-sJnWLe7KvyNkdQ@free-tier6.gcp-asia-southeast1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&options=--cluster%3Dwide-ogress-2468"
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
