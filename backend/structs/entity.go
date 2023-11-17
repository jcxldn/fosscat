package structs

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID  `gorm:"primaryKey;type:uuid;column:entityId"`
	Checkouts []Checkout `gorm:"foreignKey:checkoutId"`

	// gorm.Model (excluding ID)
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
