package util

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func IsUuidFree[T any](db *gorm.DB, id uuid.UUID, obj *T) bool {
	// NOTE: == works in SQLite, not in Postgres
	err := db.Model(obj).Select("id").Where("id = ?", id.String()).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Record not found, so user id is free
		return true
	} else {
		// Record was returned successfully, therefore the user exists
		return false
	}
}
