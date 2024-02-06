package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Model *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL not found in .env")
	}

	db, err := gorm.Open(mysql.Open(DATABASE_URL))
	if err != nil {
		panic(err)
	}

	log.Println("Database connected.")
	Model = db

	Migration()
}
