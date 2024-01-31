package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jcxldn/fosscat/backend/structs"
	"github.com/jcxldn/fosscat/backend/util"
	"gorm.io/gorm"
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

// Returns token, claims and nil error if jwt can be verified
func VerifyJwt(jwtStr string) (token *jwt.Token, claims *Claims, err error) {
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

func GetUserForJwt(db *gorm.DB, token *jwt.Token) (*structs.User, error) {
	uid, err := token.Claims.GetSubject()
	if err == nil {
		uuid, err2 := uuid.Parse(uid)
		if err2 == nil {
			// uid is a valid UUID

			user, err3 := util.GetObjectById[structs.User](db, uuid)

			if err3 == nil {
				// User found sucessfully.
				return user, nil
			} else {
				// user does not exist, or other error fetching
				return nil, errors.New("unable to fetch user from database by user id")
			}
		} else {
			// uid is not a valid UUID
			return nil, errors.New("subject was not a valid user id")
		}
	} else {
		// subject was not present in jwt
		return nil, errors.New("subject not present in jwt")
	}
}
