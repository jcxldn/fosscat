package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jcxldn/fosscat/backend/structs"
)

type Claims struct {
	jwt.RegisteredClaims
}

func getExpiryTime() time.Time {
	time := time.Now()
	time = time.AddDate(0, 0, 1) // one day
	return time
}

func NewJwt(user structs.User) (str string, err error) {
	claims := Claims{
		jwt.RegisteredClaims{
			Issuer:    "fosscat",
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(getExpiryTime()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES384, claims)

	str, err = token.SignedString(key)

	return
}

func VerifyJwt(jwtStr string, user structs.User) (token *jwt.Token, claims *Claims, err error) {
	token, err = jwt.ParseWithClaims(jwtStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return &key.PublicKey, nil
	})

	if token == nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return token, claims, nil
	} else {
		return nil, nil, err
	}
}
