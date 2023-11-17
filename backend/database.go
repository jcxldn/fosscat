package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jcxldn/fosscat/backend/structs"
)

func connect() *gorm.DB {
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

	log.Println("[database] connected. migrating...")

	// "Migrate" the schema
	// This will create tables, keys, columns, etc. Everything really.
	// See https://gorm.io/docs/migration.html
	// Note that we need to pass each struct in our schema.

	db.AutoMigrate(&structs.Checkout{})
	// Item has a dependency on Entity, so do them in the correct order
	// to avoid "relation does not exist" error during table creation.
	db.AutoMigrate(&structs.Entity{})
	db.AutoMigrate(&structs.Item{})
	db.AutoMigrate(&structs.User{})

	log.Println("[database] migrated, done.")

	return db
}
