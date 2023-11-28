package jwt

import "github.com/golang-jwt/jwt/v5"

func NewJwt() (str string, err error) {
	token := jwt.New(jwt.SigningMethodES384)

	str, err = token.SignedString(key)

	return
}
