package database_test

import (
	"testing"

	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/test/common"
	"github.com/stretchr/testify/suite"
)

type ItemTestSuite struct {
	common.EntityDatabaseTestSuite
	item *structs.Item
}

func (s *EntityTestSuite) TestCreateItemEmpty() {
	newItem := model.NewItem{}
	item, err := database.CreateItem(s.DB, newItem)

	s.Assertions.Empty(item.Entities)
	s.Assertions.Nil(err)
}

func (s *EntityTestSuite) TestCreateItemNameOnly() {
	title := "Example Item"
	newItem := model.NewItem{Title: &title}
	item, err := database.CreateItem(s.DB, newItem)

	s.Assertions.Equal(item.Title, title)
	s.Assertions.Empty(item.Entities)
	s.Assertions.Nil(err)
}

func (s *EntityTestSuite) TestCreateItemSingleEntityOnly() {
	s.CreateEntity()

	existingEntity := model.ExistingEntity{ID: s.Entity.ID.String()}
	existingEntities := []*model.ExistingEntity{}
	existingEntities = append(existingEntities, &existingEntity)

	newItem := model.NewItem{Entities: existingEntities}
	item, err := database.CreateItem(s.DB, newItem)

	s.Assertions.Equal(item.Entities[0].ID, s.Entity.ID)
	s.Assertions.Len(item.Entities, 1)
	s.Assertions.Nil(err)
}

func TestItemTestSuite(t *testing.T) {
	suite.Run(t, new(ItemTestSuite))
}
