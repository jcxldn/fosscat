package common

import (
	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
)

type UserDatabaseTestSuite struct {
	// inherit from DatabaseTestSuite
	DatabaseTestSuite
	User *structs.User
}

func (s *UserDatabaseTestSuite) CreateUser() {
	newUser := model.NewUser{FirstName: "Example", LastName: "User", Email: "user@example.com"}
	user, err := database.CreateUser(s.DB, newUser)

	if err != nil {
		panic(err)
	}

	s.User = user

	s.Assertions.Equal(user.FirstName, newUser.FirstName)
	s.Assertions.Equal(user.LastName, newUser.LastName)
	s.Assertions.Equal(user.Email, newUser.Email)
}
