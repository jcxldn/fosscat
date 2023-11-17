package structs

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Checkout struct {
	gorm.Model
	ID         uuid.UUID
	User       User `gorm:"foreignKey:ID"`
	TakeDate   time.Time
	ReturnDate time.Time
}
