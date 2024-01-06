package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB, input model.NewItem) (*structs.Item, error) {
	// Create a Item struct
	item := structs.Item{}

	isFreeUuid := false
	for !isFreeUuid {
		// Generate a UUID for the item id
		item.ID = uuid.New()
		// Check that the UUID has not been used already
		// If true, will break out of this for loop and continue
		isFreeUuid = util.IsUuidFree[structs.Item](db, item.ID)
	}

	// When we get here, we have found a non-used UUID.

	// MAJOR STEP: Assign any given entities to the item

	// Check if we were given any entity IDs and iterate over them
	// First variable is index, but we are not using it so put _. (standard practice)
	for _, entityId := range input.Entities {
		// Attempt to parse this uuid
		id, parseErr := uuid.Parse(entityId.ID)

		if parseErr == nil {
			// UUID is valid, attempt to get a object by that ID
			entity, dbErr := util.GetObjectById[structs.Entity](db, id)

			if dbErr == nil {
				// Object exists
				item.Entities = append(item.Entities, *entity)
			} else {
				// Object does not exist
				// Let's cancel the whole transaction.
				// TODO: return problematic UUID
				return nil, errors.New("unable to create item: entity does not exist")
			}
		} else {
			// entityId is not a valid UUID
			// Let's cancel the whole transaction.
			// TODO: return problematic UUID
			return nil, errors.New("unable to create item: unable to parse entity id")
		}
	}

	// All entities have been parsed and added successfully
	// If error, transaction would have been cancelled.

	// Create the db entry.
	db.Create(&item)

	// Return the result
	return &item, nil
}
