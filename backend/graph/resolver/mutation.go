package resolver

import (
	"context"

	"github.com/jcxldn/fosscat/backend/authResolver"
	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util/jwt"
	"github.com/jcxldn/fosscat/backend/util/login"
	"golang.org/x/crypto/bcrypt"
)

// CreateCheckout is the resolver for the createCheckout field.
func (r *mutationResolver) CreateCheckout(ctx context.Context, input model.NewCheckout) (*structs.Checkout, error) {
	// Use the mutation resolver to make this an authenticated route
	return authResolver.Resolver[*structs.Checkout](ctx, func(user *structs.User) authResolver.ReturnFactory {
		return authResolver.Return(database.CreateCheckout(r.DB, input))
	})
}

// CreateEntity is the resolver for the createEntity field.
func (r *mutationResolver) CreateEntity(ctx context.Context) (*structs.Entity, error) {
	// Use the mutation resolver to make this an authenticated route
	return authResolver.Resolver[*structs.Entity](ctx, func(user *structs.User) authResolver.ReturnFactory {
		return authResolver.Return(database.CreateEntity(r.DB))
	})
}

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*structs.Item, error) {
	// Use the mutation resolver to make this an authenticated route
	return authResolver.Resolver[*structs.Item](ctx, func(user *structs.User) authResolver.ReturnFactory {
		return authResolver.Return(database.CreateItem(r.DB, input))
	})
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*structs.User, error) {
	return database.CreateUser(r.DB, input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.LoginResponse, error) {
	// Create a user struct
	user := structs.User{Email: email}

	// Fetch the user from the database
	// SELECT * FROM users ORDER BY id LIMIT 1;
	// validation for only one uid per email needs to be done in createUser.
	r.DB.First(&user)

	res := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))

	// If res == nil, password hash matches and so shouldLogin is true. A jwt response is returned.
	// Else, shouldLogin is false and an error response is given
	return login.HandleLoginActions(res == nil, &user, false)
}

// LoginRefresh is the resolver for the loginRefresh field.
func (r *mutationResolver) LoginRefresh(ctx context.Context, refresh string) (*model.LoginResponse, error) {
	// Verify that we signed the given jwt (required argument)
	token, _, err := jwt.VerifyJwt(refresh)
	if err == nil {
		// JWT was verified sucessfully (we created it)
		// Let's fetch the user for this token
		user, err2 := jwt.GetUserForJwt(r.DB, token)

		// If err2 == nil, user was fetched sucessfully
		// Else, user is nil (this is okay as false shouldLogin does not call *user)
		return login.HandleLoginActions(err2 == nil, user, false)
	} else {
		// Return error handle login action (user var is not used in false shouldLogin)
		return login.HandleLoginActions(false, nil, false)
	}
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
