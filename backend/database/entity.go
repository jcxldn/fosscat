package database

import (
	"github.com/google/uuid"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util"
	"gorm.io/gorm"
)

func CreateEntity(db *gorm.DB) (*structs.Entity, error) {
	// Create a Entity struct
	entity := structs.Entity{}

	isFreeUuid := false
	for !isFreeUuid {
		// Generate a UUID for the user id.
		entity.ID = uuid.New()
		// Check that the UUID has not been used already
		// If true, it will break out of this for loop and continue.
		isFreeUuid = util.IsUuidFree[structs.Entity](db, entity.ID)
	}

	db.Create(&entity)

	return &entity, nil
}
