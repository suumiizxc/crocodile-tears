package config

import (
	"os"

	"github.com/suumiizxc/car-marketplace/models/client"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	// if err := godotenv.Load(".env"); err != nil {
	// 	panic(fmt.Sprintf("Failed env : %v", err))
	// }
	// dsn := os.Getenv("DB")
	dsn := "postgresql://suumii:Mongol123zxc@cmdm-instance.catbncexslfa.us-east-1.rds.amazonaws.com:5432/cmtestdb"
	connstring := os.ExpandEnv(dsn)
	database, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(
		&client.Client{},
		&client.Permission{},
		&client.ClientActivation{},
	)
	DB = database
}
