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

// An array of UploadedFile will be returned following a successful call to /upload (non-graphql)
type UploadedFile struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
