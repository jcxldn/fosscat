package structs

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Entity struct {
	gorm.Model
	ID        uuid.UUID
	Checkouts []Checkout `gorm:"foreignKey:ID"`
}
