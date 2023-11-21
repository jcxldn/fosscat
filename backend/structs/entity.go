package structs

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

// Entity has many Checkouts, Checkout.EntityID is the foreign key
type Entity struct {
	gorm.Model
	ID        uuid.UUID
	ItemID    uuid.UUID // Foreign key for Item
	Checkouts []Checkout
}
