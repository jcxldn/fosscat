package structs

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

// Item has many Entities, Entity.ItemID is the foreign key
type Item struct {
	gorm.Model
	ID       uuid.UUID
	Title    string
	Entities []Entity
}
