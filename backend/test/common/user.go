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

func (s *UserDatabaseTestSuite) TestCreateUser() {
	newUser := model.NewUser{FirstName: "Example", LastName: "User", Email: "example.user@example.com"}
	user, err := database.CreateUser(s.DB, newUser)

	if err != nil {
		panic(err)
	}

	s.User = user
}
