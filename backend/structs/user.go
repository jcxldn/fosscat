package structs

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;column:userId"`
	FirstName string
	LastName  string
	Email     string
	Hash      string

	// gorm.Model (excluding ID)
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
