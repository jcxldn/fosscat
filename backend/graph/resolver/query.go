package resolver

import (
	"context"
	"errors"

	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/structs"
	"gorm.io/gorm/clause"
)

// Checkout is the resolver for the checkout field.
func (r *queryResolver) Checkout(ctx context.Context) ([]*structs.Checkout, error) {
	// Get all checkouts. Proof of concept only, returns all fields!
	checkouts := []*structs.Checkout{}
	// Preload all associations (https://gorm.io/docs/preload.html#Preload-All)
	// TODO much later: don't preload private fields like hash unless user = logged in user.
	result := r.DB.Preload(clause.Associations).Find(&checkouts)
	return checkouts, result.Error
}

// Entity is the resolver for the entity field.
func (r *queryResolver) Entity(ctx context.Context) ([]*structs.Entity, error) {
	// Get all entities. Proof of concept only, returns all fields!
	entities := []*structs.Entity{}
	result := r.DB.Find(&entities)
	return entities, result.Error
}

// File is the resolver for the file field.
func (r *queryResolver) File(ctx context.Context) ([]*structs.File, error) {
	// Get all files. Proof of concept only, returns all fields!
	files := []*structs.File{}
	// Preload all associations (https://gorm.io/docs/preload.html#Preload-All)
	// TODO much later: don't preload private fields like hash unless user = logged in user.
	result := r.DB.Preload(clause.Associations).Find(&files)
	return files, result.Error
}

// Item is the resolver for the item field.
func (r *queryResolver) Item(ctx context.Context) ([]*structs.Item, error) {
	// Get all items. Proof of concept only, returns all fields!
	items := []*structs.Item{}
	result := r.DB.Find(&items)
	return items, result.Error
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*structs.User, error) {
	// Get all users. Proof of concept only, returns all fields!
	users := []*structs.User{}
	result := r.DB.Find(&users)
	return users, result.Error
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*structs.User, error) {
	// Check if the user context is present
	// (is present when a valid jwt is placed in the authorization header)
	user := ctx.Value("user")
	if user != nil {
		// Type assertion that "user" ctx is not nil and is of type structs.User
		userStruct := user.(*structs.User)

		// Get a copy of the struct with "private" fields (like password hash) removed.
		publicUserStruct := userStruct.GetPublicFields()

		// Return the public fields only struct.
		return &publicUserStruct, nil
	}
	return nil, errors.New("route requires authorization")
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
