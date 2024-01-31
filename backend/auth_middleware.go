package main

import (
	"context"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util/jwt"
	"gorm.io/gorm"
)

// Middleware to decode the jwt and place the associated user data into the context
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *structs.User

		// Check if there is a jwt present in the Authorization header
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader != "" {
			// Header is present.
			splitHeader := strings.Split(authHeader, "Bearer ")
			if len(splitHeader) == 2 {
				if splitHeader[1] != "" {
					tokenStr := splitHeader[1]

					token, _, err := jwt.VerifyJwt(tokenStr)
					if err == nil {
						// JWT was verified successfully.
						// Let's fetch the user for this token.
						userRes, err2 := jwt.GetUserForJwt(db, token)
						if err2 == nil {
							// User was fetched successfully
							user = userRes

						} else {
							// JWT was valid but user did not exist.
							log.Fatal("Valid JWT presented but user could not be fetched from database!")
						}
					}
				} else {
					// Authorization header contains "Bearer " but no token, ignore.
				}
			} else {
				// Auth header is invalid (doesn't contain "Bearer ")
			}
		} else {
			// Authorization header is not present, ignore.
		}

		// If user is present, set context
		if user != nil {
			// Get a copy of the parent context with the user key set
			ctx := context.WithValue(c.Request.Context(), "user", user)
			// Create a new request with the new context
			newRequest := c.Request.WithContext(ctx)
			// Overwrite the original request with the new one
			c.Request = newRequest
		}

		// Call the next handler (including the target itself)
		c.Next()

		// Request has been executed
	}
}
