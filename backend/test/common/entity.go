package common

import (
	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/structs"
)

type EntityDatabaseTestSuite struct {
	// inherit from DatabaseTestSuite
	DatabaseTestSuite
	Entity *structs.Entity
}

func (s *EntityDatabaseTestSuite) CreateEntity() {
	entity, err := database.CreateEntity(s.DB)

	if err != nil {
		panic(err)
	}

	s.Entity = entity

	s.Assertions.Nil(err)
}
