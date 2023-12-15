package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jcxldn/fosscat/backend/structs"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	// Attempt to connect to the db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Migrate(db)

	return db
}

func Migrate(db *gorm.DB) {
	log.Println("[database] migrating...")

	var errors [4]error

	// "Migrate" the schema
	// This will create tables, keys, columns, etc. Everything really.
	// See https://gorm.io/docs/migration.html
	// Note that we need to pass each struct in our schema.

	errors[0] = db.AutoMigrate(&structs.Checkout{})
	// Item has a dependency on Entity, so do them in the correct order
	// to avoid "relation does not exist" error during table creation.
	errors[1] = db.AutoMigrate(&structs.Entity{})
	errors[2] = db.AutoMigrate(&structs.Item{})
	errors[3] = db.AutoMigrate(&structs.User{})

	for index, err := range errors {
		if err != nil {
			log.Panicf("Error with migrate call %d: '%s'", index, err)
		}
	}

	log.Println("[database] migrated, done.")
}
