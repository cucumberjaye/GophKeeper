package tokens

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const signKey = "test"

type authClaims struct {
	*jwt.StandardClaims

	UserID string `json:"userID"`
}

func CreateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &authClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		userID,
	})

	return token.SignedString([]byte(signKey))
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("parse jwt token failed with error: %w", err)
	}

	claims, ok := token.Claims.(*authClaims)
	if !ok {
		return "", errors.New("token claims are not of type *authClaims")
	}

	return claims.UserID, nil
}
