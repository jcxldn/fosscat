package authResolver

import (
	"context"
	"errors"

	"github.com/jcxldn/fosscat/backend/structs"
)

type ReturnFactory func() (any, error)

func Return(a any, b error) ReturnFactory {
	return func() (any, error) {
		return a, b
	}
}

func QueryResolver[T any](ctx context.Context, resolver func(user *structs.User) ReturnFactory) (T, error) {
	// Check if the user context is present
	// (is present when a valid jwt is placed in the authorization header)
	user := ctx.Value("user")
	if user != nil {
		// Type assertion that "user" ctx is not nil and is of type structs.User
		userStruct := user.(*structs.User)

		returnData := resolver(userStruct)
		data, err := returnData()
		return data.(T), err
	} else {
		var null T
		return null, errors.New("route requires authorization")
	}
}
