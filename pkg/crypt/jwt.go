package crypt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewJwtToken(secretKey []byte, issuer string, lifetime int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(lifetime) * time.Second)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		})
	return t.SignedString(secretKey)
}

func VerifyToken(tokenString, secretKey string) (*jwt.Token, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return token, nil
}
