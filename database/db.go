package database

import (
	"bkndOpenMind/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(config.GetDBurl()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
