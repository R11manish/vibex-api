package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dbUser := os.Getenv("VIBEX_PG_USER")
	dbPass := os.Getenv("VIBEX_PG_PWD")
	dbName := os.Getenv("VIBEX_PG_DB")
	dbPort := os.Getenv("VIBEX_PG_PORT")

	// Create the connection string
	connStr := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbPort)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the PostgreSQL database")
	return db, nil
}
