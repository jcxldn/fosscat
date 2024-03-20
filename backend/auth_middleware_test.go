package main_test

import (
	"testing"

	main "github.com/jcxldn/fosscat/backend"
	"github.com/jcxldn/fosscat/backend/test/common"
	"github.com/jcxldn/fosscat/backend/util/jwt"
	"github.com/stretchr/testify/suite"
)

type AuthMiddlewareTestSuite struct {
	// Inherit from GinTestSuite
	common.GinTestSuite
}

func (s *AuthMiddlewareTestSuite) TestNoAuthorizationHeader() {
	handlerFunc := main.AuthMiddleware(s.DB)
	handlerFunc(s.Context)
	s.Assertions.Equal(main.RESULT_NO_AUTH_HEADER, s.Context.Request.Context().Value("result"))
}

func (s *AuthMiddlewareTestSuite) TestInvalidAuthorizationHeader() {
	handlerFunc := main.AuthMiddleware(s.DB)
	s.Context.Request.Header.Add("Authorization", "Basic supersecure")
	handlerFunc(s.Context)
	s.Assertions.Equal(main.RESULT_INVALID_AUTH_HEADER, s.Context.Request.Context().Value("result"))
}
func (s *AuthMiddlewareTestSuite) TestNoTokenInAuthorizationHeader() {
	handlerFunc := main.AuthMiddleware(s.DB)
	s.Context.Request.Header.Add("Authorization", "Bearer ")
	handlerFunc(s.Context)
	s.Assertions.Equal(main.RESULT_NO_TOKEN_IN_AUTH_HEADER, s.Context.Request.Context().Value("result"))
}

func (s *AuthMiddlewareTestSuite) TestInvalidTokenInAuthorizationHeader() {
	handlerFunc := main.AuthMiddleware(s.DB)
	s.Context.Request.Header.Add("Authorization", "Bearer hello")
	handlerFunc(s.Context)
	s.Assertions.Equal(main.RESULT_INVALID_JWT, s.Context.Request.Context().Value("result"))
}

func (s *AuthMiddlewareTestSuite) TestSuccessful() {
	s.CreateUser()
	jwt, err := jwt.NewJwt(*s.User)
	s.Assertions.Nil(err)
	// s.user is the new user

	handlerFunc := main.AuthMiddleware(s.DB)
	s.Context.Request.Header.Add("Authorization", "Bearer "+jwt)
	handlerFunc(s.Context)
	s.Assertions.Equal(main.RESULT_USER_FETCHED_SUCESSFULLY, s.Context.Request.Context().Value("result"))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to s.Run
func TestAuthMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareTestSuite))
}
