package database

import (
	"bkndOpenMind/config"
	"bkndOpenMind/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(config.GetDBurl()), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to conncet\n", err.Error())
		os.Exit(2)
		// panic("failed to connect database")
	}

	log.Println("Connections is succed")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migration")

	// add migrations
	db.AutoMigrate(&models.User{}, &models.Order{}, &models.Product{})

	DB = DbInstance{Db: db}
}
