package structs

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Item struct {
	gorm.Model
	ID       uuid.UUID
	Title    string
	Entities []Entity `gorm:"foreignKey:ID"`
}
