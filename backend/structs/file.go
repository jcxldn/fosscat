package structs

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

// File belongs to a User, Uploader.ID (UploaderID) is the foreign key
type File struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid"`
	Uploader   User
	UploaderID uuid.UUID `gorm:"type:uuid"`
	Name       string
	Type       FileTypes
	Data       []byte
}
