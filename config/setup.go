package config

import (
	"os"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	// if err := godotenv.Load(".env"); err != nil {
	// 	panic(fmt.Sprintf("Failed env : %v", err))
	// }
	// dsn := os.Getenv("DB")
	dsn := "postgresql://suumiizxc:HNjLiCKWrjhz3BQE3myiZg@free-tier6.gcp-asia-southeast1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&options=--cluster%3Dsteely-runner-2815"
	connstring := os.ExpandEnv(dsn)
	database, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	// database.AutoMigrate(
	// 	&client.Client{},
	// 	&client.Permission{},

	// 	&marketplace.Car{},
	// 	&marketplace.CarImage{},
	// 	&marketplace.DiagonisImage{},
	// )

	DB = database
}
