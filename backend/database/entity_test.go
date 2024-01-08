package database_test

import (
	"testing"

	"github.com/jcxldn/fosscat/backend/test/common"
	"github.com/stretchr/testify/suite"
)

type EntityTestSuite struct {
	common.EntityDatabaseTestSuite
}

func (s *EntityTestSuite) TestCreateEntity() {
	s.CreateEntity()
}

func TestEntityTestSuite(t *testing.T) {
	suite.Run(t, new(EntityTestSuite))
}
