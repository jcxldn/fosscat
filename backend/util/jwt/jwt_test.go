package jwt_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jcxldn/fosscat/backend/test/common"
	"github.com/jcxldn/fosscat/backend/util/jwt"
	"github.com/stretchr/testify/suite"
)

type UtilJwtTestSuite struct {
	common.UserDatabaseTestSuite
	jwt          *string
	tempFilePath string
}

func (s *UtilJwtTestSuite) SetupSuite() {
	s.UserDatabaseTestSuite.SetupSuite() // Call parent SetupSuite
	s.CreateUser()                       // Create a user for this test suite run

	tempFilePath := fmt.Sprintf("%s/fosscat-test.pem", os.TempDir())

	fmt.Printf("Using temporary file '%s'", tempFilePath)

	s.tempFilePath = tempFilePath

	os.Setenv("JWT_PK_FILE", tempFilePath)
	jwt.SetupKey()
}

func (s *UtilJwtTestSuite) TeardownSuite() {
	fmt.Printf("Removing temporary file '%s'", s.tempFilePath)
}
func (s *UtilJwtTestSuite) TestCreateJwt() {
	token, err := jwt.NewJwt(*s.User)

	s.Assertions.Nil(err)

	s.jwt = &token
}

func (s *UtilJwtTestSuite) TestVerifyJwtValid() {
	token, err := jwt.NewJwt(*s.User)

	s.Assertions.Nil(err)

	parsedToken, claims, err2 := jwt.VerifyJwt(token, *s.User)

	s.Assertions.Nil(err2)

	s.Assertions.NotNil(parsedToken)
	s.Assertions.True(parsedToken.Valid)

	s.Assertions.Equal(claims.Subject, s.User.ID.String())

}

func (s *UtilJwtTestSuite) TestVerifyJwtInvalid() {

	parsedToken, claims, err := jwt.VerifyJwt("invalid jwt", *s.User)

	s.Assertions.NotNil(err)
	s.Assertions.Nil(parsedToken)
	s.Assertions.Nil(claims)
}

func TestUtilJwtTestSuite(t *testing.T) {
	suite.Run(t, new(UtilJwtTestSuite))
}
