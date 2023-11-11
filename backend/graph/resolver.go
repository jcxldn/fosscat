package graph

import (
	"log"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *gorm.DB
}

// Set DB object if not already set
func (r *Resolver) UpdateDb(db *gorm.DB) {
	if r.db == nil {
		log.Printf("[resolver] db not set, setting.")
		r.db = db
	} else {
		log.Printf("[resolver] WARN: db is already set, not overwriting.")
	}
}
