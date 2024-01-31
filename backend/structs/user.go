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

func (u User) GetPublicFields() User {
	publicUser := User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
	return publicUser
}
