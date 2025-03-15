package config

import (
	"sawittree/models"

	"gorm.io/driver/postgres"

	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {

	connectionString := "user=postgres password=root dbname=sawittree host=localhost port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	InitMigrate()
}

func InitDBTest() {
	connectionString := "user=postgres password=root dbname=sawittree host=localhost port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Estate{}, &models.Tree{})
}
