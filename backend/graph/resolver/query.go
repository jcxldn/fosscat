package resolver

import (
	"context"

	"github.com/jcxldn/fosscat/backend/authResolver"
	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/structs"
	"gorm.io/gorm/clause"
)

// Checkout is the resolver for the checkout field.
func (r *queryResolver) Checkout(ctx context.Context) ([]*structs.Checkout, error) {
	// Use the query resolver to make this an authenticated route.
	// Even though user is not used, it must still be passed to the func.
	return authResolver.QueryResolver[[]*structs.Checkout](ctx, func(user *structs.User) authResolver.ReturnFactory {
		// Get all entities, returns all fields!
		checkouts := []*structs.Checkout{}
		// Preload all associations (https://gorm.io/docs/preload.html#Preload-All)
		// TODO much later: don't preload private fields like hash unless user = logged in user.
		result := r.DB.Preload(clause.Associations).Find(&checkouts)
		return authResolver.Return(checkouts, result.Error)
	})
}

// Entity is the resolver for the entity field.
func (r *queryResolver) Entity(ctx context.Context) ([]*structs.Entity, error) {
	// Use the query resolver to make this an authenticated route.
	// Even though user is not used, it must still be passed to the func.
	return authResolver.QueryResolver[[]*structs.Entity](ctx, func(user *structs.User) authResolver.ReturnFactory {
		// Get all entities, returns all fields!
		entities := []*structs.Entity{}
		result := r.DB.Find(&entities)
		return authResolver.Return(entities, result.Error)
	})
}

// File is the resolver for the file field.
func (r *queryResolver) File(ctx context.Context) ([]*structs.File, error) {
	// Use the query resolver to make this an authenticated route.
	// Even though user is not used, it must still be passed to the func.
	return authResolver.QueryResolver[[]*structs.File](ctx, func(user *structs.User) authResolver.ReturnFactory {
		// Get all files, returns all fields!
		files := []*structs.File{}
		// Preload all associations (https://gorm.io/docs/preload.html#Preload-All)
		// TODO much later: don't preload private fields like hash unless user = logged in user.
		result := r.DB.Preload(clause.Associations).Find(&files)
		return authResolver.Return(files, result.Error)
	})
}

// Item is the resolver for the item field.
func (r *queryResolver) Item(ctx context.Context) ([]*structs.Item, error) {
	// Use the query resolver to make this an authenticated route.
	// Even though user is not used, it must still be passed to the func.
	return authResolver.QueryResolver[[]*structs.Item](ctx, func(user *structs.User) authResolver.ReturnFactory {
		// Get all items, returns all fields!
		items := []*structs.Item{}
		result := r.DB.Find(&items)
		return authResolver.Return(items, result.Error)
	})
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*structs.User, error) {
	// Use the query resolver. The anonymous function passed as a parameter will be called if auth context was set.
	// If auth context was not set, query resolver will return an error message.
	return authResolver.QueryResolver[[]*structs.User](ctx, func(user *structs.User) authResolver.ReturnFactory {
		// Create an empty array of user pointers
		users := []*structs.User{}
		// Populate the array
		result := r.DB.Find(&users)

		// Convert return type to generic ReturnFactory
		// If error occurred, users will be nil.
		// If no error occurred, result.Error will be nil. No if/else block needed; can be done in one line.
		return authResolver.Return(users, result.Error)
	})
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*structs.User, error) {
	// Use the query resolver. The anonymous function passed as a parameter will be called if auth context was set.
	// If auth context was not set, query resolver will return an error message.
	return authResolver.QueryResolver[*structs.User](ctx, func(user *structs.User) authResolver.ReturnFactory {
		// Return the user object from context.
		return authResolver.Return(user, nil)
	})
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
