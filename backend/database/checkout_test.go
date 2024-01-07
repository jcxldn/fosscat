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

func (s *CheckoutTestSuite) TestCreateCheckoutAllFields() {
	existingUser := model.ExistingUser{ID: s.User.ID.String()}
	takeDate := time.Time{}
	returnDate := time.Time{}.AddDate(0, 0, 1)
	newCheckout := model.NewCheckout{User: &existingUser, TakeDate: &takeDate, ReturnDate: &returnDate}

	checkout, err := database.CreateCheckout(s.DB, newCheckout)

	s.Assertions.NotNil(checkout.TakeDate)
	s.Assertions.NotNil(checkout.ReturnDate)
	s.Assertions.Equal(checkout.User.ID, s.User.ID)
	s.Assertions.Equal(checkout.UserID, s.User.ID)

	s.Assertions.Nil(err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to s.Run
func TestCheckoutTestSuite(t *testing.T) {
	suite.Run(t, new(CheckoutTestSuite))
}
