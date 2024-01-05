package util_test

import (
	"testing"

	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/test/common"
	"github.com/jcxldn/fosscat/backend/util"
	"github.com/stretchr/testify/suite"
)

type UtilGetTestSuite struct {
	common.DatabaseTestSuite
	user *structs.User
}

func (s *UtilGetTestSuite) TestCreateUser() {
	newUser := model.NewUser{FirstName: "Example", LastName: "User", Email: "example.user@example.com"}
	user, err := database.CreateUser(s.DB, newUser)

	if err != nil {
		panic(err)
	}

	s.user = user
}

func (s *UtilGetTestSuite) TestCreateUserInvalidEmail() {
	newUser := model.NewUser{FirstName: "Example", LastName: "User", Email: "example*email"}
	_, err := database.CreateUser(s.DB, newUser)

	s.Assertions.EqualError(err, "email address not valid")
}

func (s *UtilGetTestSuite) TestCreateUserInvalidPasswordTooLong() {
	// >72 byte password is too long (72 byte limit)
	newUser := model.NewUser{Email: "example.user@example.com", Password: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc tincidunt. "}
	_, err := database.CreateUser(s.DB, newUser)

	s.Assertions.EqualError(err, "password does not satisfy requirements")
}

func (s *UtilGetTestSuite) TestGetObjectByIdUser() {
	user, err := util.GetObjectById[structs.User](s.DB, s.user.ID)

	if err != nil {
		panic(err)
	}

	s.Assertions.Equal(s.user.ID, user.ID)
}

func (s *UtilGetTestSuite) TestUuidIsFreeUser() {
	s.Assertions.False(util.IsUuidFree[structs.User](s.DB, s.user.ID))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to s.Run
func TestUtilGetTestSuite(t *testing.T) {
	suite.Run(t, new(UtilGetTestSuite))
}
