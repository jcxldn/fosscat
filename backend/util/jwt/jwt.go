package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jcxldn/fosscat/backend/structs"
)

func getExpiryTime() int64 {
	time := time.Now()
	time = time.AddDate(0, 0, 1) // one day
	return time.Unix()
}

func NewJwt(user structs.User) (str string, err error) {
	claims := jwt.MapClaims{
		"iss": "fosscat",
		"sub": user.ID,
		"exp": getExpiryTime(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES384, claims)

	str, err = token.SignedString(key)

	return
}
