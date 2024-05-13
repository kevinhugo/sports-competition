package helpers

import (
	"sports-competition/app/env"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AccessToken struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateAccessToken(ID string, username string) (string, error) {

	claims := &AccessToken{
		ID:       ID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(env.ACCESS_TOKEN_EXPIRED_IN_MINUTES) * time.Hour).Unix(),
			// IssuedAt:  time.Now().Unix(),
			// NotBefore: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(env.ACCESS_TOKEN_SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
