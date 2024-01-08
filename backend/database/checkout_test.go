package database_test

import (
	"testing"
	"time"

	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/test/common"
	"github.com/stretchr/testify/suite"
)

type CheckoutTestSuite struct {
	common.UserDatabaseTestSuite
}

func (s *CheckoutTestSuite) SetupSuite() {
	s.UserDatabaseTestSuite.SetupSuite() // Call parent SetupSuite
	s.CreateUser()                       // Create a user for this test suite run
}

func (s *CheckoutTestSuite) TestCreateCheckoutAllFields() {
	existingUser := model.ExistingUser{ID: s.User.ID.String()}
	takeDate := time.Now()
	returnDate := time.Now().AddDate(0, 0, 1)
	newCheckout := model.NewCheckout{User: &existingUser, TakeDate: &takeDate, ReturnDate: &returnDate}

	checkout, err := database.CreateCheckout(s.DB, newCheckout)

	s.Assertions.False(checkout.TakeDate.IsZero())
	s.Assertions.False(checkout.ReturnDate.IsZero())

	s.Assertions.Equal(checkout.User.ID, s.User.ID)
	s.Assertions.Equal(checkout.UserID, s.User.ID)

	s.Assertions.Nil(err)
}

func (s *CheckoutTestSuite) TestCreateCheckoutNoReturnDate() {

	existingUser := model.ExistingUser{ID: s.User.ID.String()}
	takeDate := time.Now()
	newCheckout := model.NewCheckout{User: &existingUser, TakeDate: &takeDate, ReturnDate: nil}

	checkout, err := database.CreateCheckout(s.DB, newCheckout)

	s.Assertions.False(checkout.TakeDate.IsZero())
	s.Assertions.True(checkout.ReturnDate.IsZero())

	s.Assertions.Equal(checkout.User.ID, s.User.ID)
	s.Assertions.Equal(checkout.UserID, s.User.ID)

	s.Assertions.Nil(err)
}

func (s *CheckoutTestSuite) TestCreateCheckoutNoTakeDate() {

	existingUser := model.ExistingUser{ID: s.User.ID.String()}
	returnDate := time.Now().AddDate(0, 0, 1)
	newCheckout := model.NewCheckout{User: &existingUser, TakeDate: nil, ReturnDate: &returnDate}

	checkout, err := database.CreateCheckout(s.DB, newCheckout)

	s.Assertions.True(checkout.TakeDate.IsZero())
	s.Assertions.False(checkout.ReturnDate.IsZero())

	s.Assertions.Equal(checkout.User.ID, s.User.ID)
	s.Assertions.Equal(checkout.UserID, s.User.ID)

	s.Assertions.Nil(err)
}

func (s *CheckoutTestSuite) TestCreateCheckoutOnlyID() {

	existingUser := model.ExistingUser{ID: s.User.ID.String()}
	newCheckout := model.NewCheckout{User: &existingUser, TakeDate: nil, ReturnDate: nil}

	checkout, err := database.CreateCheckout(s.DB, newCheckout)

	s.Assertions.True(checkout.TakeDate.IsZero())
	s.Assertions.True(checkout.ReturnDate.IsZero())

	s.Assertions.Equal(checkout.User.ID, s.User.ID)
	s.Assertions.Equal(checkout.UserID, s.User.ID)

	s.Assertions.Nil(err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to s.Run
func TestCheckoutTestSuite(t *testing.T) {
	suite.Run(t, new(CheckoutTestSuite))
}
