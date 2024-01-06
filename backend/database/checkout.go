package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util"
	"gorm.io/gorm"
)

func CreateCheckout(db *gorm.DB, input model.NewCheckout) (*structs.Checkout, error) {
	// Create a Checkout struct
	checkout := structs.Checkout{}

	isFreeUuid := false
	for !isFreeUuid {
		// Generate a UUID for the checkout id.
		checkout.ID = uuid.New()
		// Check that the UUID has not been used already
		// If true, it will break out of this for loop and continue.
		isFreeUuid = util.IsUuidFree[structs.Checkout](db, checkout.ID)
	}

	// When we get here, we have found a non-used UUID.

	// Assign the checkout to a user.
	id, parseErr := uuid.Parse(input.User.ID)

	if parseErr == nil {

		user, dbErr := util.GetObjectById[structs.User](db, id)

		if dbErr == nil {
			checkout.User = *user

			// Create the database entry
			db.Create(&checkout)

			return &checkout, nil
		} else {
			return nil, errors.New("unable to create checkout: user does not exist")
		}
	} else {
		return nil, errors.New("unable to create checkout: unable to parse user id")
	}
}
