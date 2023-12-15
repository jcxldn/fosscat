package util_test

import (
	"fmt"
	"testing"

	"github.com/jcxldn/fosscat/backend/test/common"
	"github.com/stretchr/testify/suite"
)

type UtilGetTestSuite struct {
	common.DatabaseTestSuite
}

func (s *UtilGetTestSuite) TestGetObjectByIdUser() {
	fmt.Println("uwu")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to s.Run
func TestUtilGetTestSuite(t *testing.T) {
	suite.Run(t, new(UtilGetTestSuite))
}
