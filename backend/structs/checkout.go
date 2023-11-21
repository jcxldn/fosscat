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
	User       User
	UserID     uuid.UUID
	TakeDate   time.Time
	ReturnDate time.Time
}
