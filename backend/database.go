package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/jcxldn/fosscat/backend/db/structs"
)

func connect() *gorm.DB {
	// Attempt to connect to the db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("[database] connected. migrating...")

	// "Migrate" the schema
	// This will create tables, keys, columns, etc. Everything really.
	// See https://gorm.io/docs/migration.html
	// Note that we need to pass each struct in our schema.
	db.AutoMigrate(&structs.User{})

	log.Println("[database] migrated, done.")

	return db
}
