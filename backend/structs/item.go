package structs

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Item struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;column:itemId"`
	Title    string
	Entities []Entity `gorm:"foreignKey:entityId"`

	// gorm.Model (excluding ID)
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
