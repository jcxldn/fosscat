package database_test

import (
	"testing"

	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/test/common"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	common.DatabaseTestSuite
	user *structs.User
}

func (s *UserTestSuite) TestCreateUser() {
	newUser := model.NewUser{FirstName: "Example", LastName: "User", Email: "example.user@example.com"}
	user, err := database.CreateUser(s.DB, newUser)

	if err != nil {
		panic(err)
	}

	s.user = user
}

func (s *UserTestSuite) TestCreateUserInvalidEmail() {
	newUser := model.NewUser{FirstName: "Example", LastName: "User", Email: "example*email"}
	_, err := database.CreateUser(s.DB, newUser)

	s.Assertions.EqualError(err, "email address not valid")
}

func (s *UserTestSuite) TestCreateUserInvalidPasswordTooLong() {
	// >72 byte password is too long (72 byte limit)
	newUser := model.NewUser{Email: "example.user@example.com", Password: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc tincidunt. "}
	_, err := database.CreateUser(s.DB, newUser)

	s.Assertions.EqualError(err, "password does not satisfy requirements")
}

func (s *UserTestSuite) TestGetAllUsers() {
	users, err := database.GetAllUsers(s.DB)

	// Assert that there was no error
	s.Assertions.Empty(err)
	// Check that the user created in this test is in the response
	s.Assertions.Len(users, 1)
	s.Assertions.Equal(users[0].ID, s.user.ID)

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to s.Run
func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
