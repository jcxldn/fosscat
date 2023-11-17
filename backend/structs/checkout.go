package structs

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Checkout struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;column:checkoutId"`
	User       User      `gorm:"foreignKey:userId"`
	TakeDate   time.Time
	ReturnDate time.Time

	// gorm.Model (excluding ID)
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
