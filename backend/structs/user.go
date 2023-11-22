package structs

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid"`
	FirstName string
	LastName  string
	Email     string
	Hash      string
}
