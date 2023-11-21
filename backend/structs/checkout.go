package structs

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// Checkout belongs to a User, User.ID (UserID) is the foreign key
type Checkout struct {
	gorm.Model
	ID         uuid.UUID
	EntityID   uuid.UUID // Foreign key for Entity
	User       User
	UserID     uuid.UUID
	TakeDate   time.Time
	ReturnDate time.Time
}
