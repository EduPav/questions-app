package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Create variables in a global context to use it in other files of my app.
var DSN = "host=localhost user=eduyio password=admin dbname=qappdb port=5432"
var DB *gorm.DB

func DBConnection() {
	var err error

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected")
	}
}
