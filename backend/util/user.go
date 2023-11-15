package util

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jcxldn/fosscat/backend/structs"
	"gorm.io/gorm"
)

func IsUuidFree(db *gorm.DB, id uuid.UUID) bool {
	// Run a SQL query to see if the user exists
	var user structs.User
	err := db.Model(&structs.User{}).Select("id").Where("id == ?", id.String()).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Record not found, so user id is free
		return true
	} else {
		// Record was returned successfully, therefore the user exists
		return false
	}
}
