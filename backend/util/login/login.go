package login

import (
	"errors"

	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util/jwt"
)

func HandleLoginActions(shouldLogin bool, user *structs.User, generateRefresh bool) (*model.LoginResponse, error) {
	if shouldLogin { // validation step passed in caller func
		// generate and return a LoginResponse containing a access jwt
		lr := model.LoginResponse{Success: true}

		var jwtStr string
		var err error

		if generateRefresh {
			jwtStr, err = jwt.NewRefreshJwt(*user)
		} else {
			jwtStr, err = jwt.NewAccessJwt(*user)
		}
		if err == nil {
			// token was created sucessfully, return it in a LoginResponse
			lr.Jwt = &jwtStr
			return &lr, nil
		} else {
			// token was not created
			lr := model.LoginResponse{Success: false}
			return &lr, errors.New("error creating jwt")
		}
	} else {
		// Validation step failed in caller func (eg. password does not match)
		// Return a failed loginresponse
		lr := model.LoginResponse{Success: false}
		return &lr, nil
	}
}
