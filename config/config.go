package config

import (
	"os"
	"sawittree/models"

	"gorm.io/driver/postgres"

	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		dbPort = "5432" 
	}

	connectionString :=
		"host=" + dbHost +
			" user=" + dbUser +
			" password=" + dbPassword +
			" dbname=" + dbName +
			" port=" + dbPort +
			" sslmode=disable TimeZone=Asia/Jakarta"

	// connectionString := "user=postgres password=root dbname=sawittree host=localhost port=5432 sslmode=disable TimeZone=Asia/Jakarta"

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
