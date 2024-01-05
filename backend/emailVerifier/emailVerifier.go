package emailVerifier

import (
	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	ev = emailverifier.NewVerifier()
)

func VerifyEmail(email string) bool {
	// Check syntax
	syntax := ev.ParseAddress(email)

	return syntax.Valid
}
