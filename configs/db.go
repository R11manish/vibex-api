package config

import (
	"fmt"
	"log"
	"os"
	models "vibex-api/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbUser := os.Getenv("VIBEX_PG_USER")
	dbPass := os.Getenv("VIBEX_PG_PWD")
	dbName := os.Getenv("VIBEX_PG_DB")
	dbPort := os.Getenv("VIBEX_PG_PORT")

	if dbUser == "" || dbPass == "" || dbName == "" || dbPort == "" {
		log.Fatal("Database environment variables are not set.")
	}

	connStr := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db

	log.Println("Successfully connected to the PostgreSQL database")

	autoMigrate()
	Seed()
}

func GetDB() *gorm.DB {
	return DB
}

func autoMigrate() {
	env := GetEnvironment()
	if env != "local" {
		return
	}
	log.Println("Auto migrating database...")
	if err := DB.AutoMigrate(&models.Status{}, &models.User{}); err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}
}
