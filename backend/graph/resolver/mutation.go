package resolver

import (
	"context"
	"errors"

	"github.com/jcxldn/fosscat/backend/database"
	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util/jwt"
	"golang.org/x/crypto/bcrypt"
)

// CreateCheckout is the resolver for the createCheckout field.
func (r *mutationResolver) CreateCheckout(ctx context.Context, input model.NewCheckout) (*structs.Checkout, error) {
	return database.CreateCheckout(r.DB, input)
}

// CreateEntity is the resolver for the createEntity field.
func (r *mutationResolver) CreateEntity(ctx context.Context) (*structs.Entity, error) {
	return database.CreateEntity(r.DB)
}

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*structs.Item, error) {
	return database.CreateItem(r.DB, input)
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

	if res == nil {
		// Password matches, generate and return a JWT
		lr := model.LoginResponse{Success: true}
		jwt, err := jwt.NewJwt(user)
		if err == nil {
			lr.Jwt = &jwt
			return &lr, nil
		} else {
			lr := model.LoginResponse{Success: false}
			return &lr, errors.New("Error creating jwt")
		}

	} else {
		// Password does not match (bcrypt returns error on failure)
		lr := model.LoginResponse{Success: false}
		return &lr, nil
	}
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
