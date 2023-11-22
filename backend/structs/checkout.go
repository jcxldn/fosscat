package structs

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// Checkout belongs to a User, User.ID (UserID) is the foreign key
type Checkout struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid"`
	EntityID   uuid.UUID `gorm:"type:uuid"` // Foreign key for Entity
	User       User
	UserID     uuid.UUID `gorm:"type:uuid"`
	TakeDate   time.Time
	ReturnDate time.Time
}
