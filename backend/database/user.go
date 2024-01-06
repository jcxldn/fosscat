package database

import (
	"errors"

	"github.com/google/uuid"
	ev "github.com/jcxldn/fosscat/backend/emailVerifier"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, input model.NewUser) (*structs.User, error) {
	// Create a User struct from the input model
	// Set fields that do not need further validation (name only)
	user := structs.User{FirstName: input.FirstName, LastName: input.LastName}

	// Attempt to validate the user email.
	isValidEmail := ev.VerifyEmail(input.Email)
	if !isValidEmail {
		return nil, errors.New("email address not valid")
	}

	// Email passed validation, set in the user struct.
	user.Email = input.Email

	isFreeUuid := false
	for !isFreeUuid {
		// Generate a UUID for the user id.
		user.ID = uuid.New()
		// Check that the UUID has not been used already
		// If true, it will break out of this for loop and continue.
		isFreeUuid = util.IsUuidFree[structs.User](db, user.ID)
	}

	// Salt and hash the provided password
	// I am currently using bCrypt, which has the function GenerateFromPasswords.
	// This generates a random salt for us and applies it to the hash.
	// I believe the hash result and salt are stored side by side.
	// Cost of '10' for now, seems like a good balance.
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		// Hashing failed (password too long/short?)
		// TODO: make errors more readable, custom errors
		return nil, errors.New("password does not satisfy requirements")
	}

	// Hashing completed successfully, set in the user struct
	// For purposes of writeup use string for now so we can remove it later.
	user.Hash = string(hash)

	db.Create(&user)
	return &structs.User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Hash: user.Hash}, nil
}

func GetAllUsers(db *gorm.DB) ([]*structs.User, error) {
	// Get all users. Proof of concept only, returns all fields!
	users := []*structs.User{}
	result := db.Find(&users)
	return users, result.Error
}
