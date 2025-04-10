package services

import (
	"cushon/consts"
	"cushon/types"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtService struct {
	jwtKey []byte
}

// NewJwtService creates a new instance of the JWT service
func NewJwtService(jwtKey []byte) (jwtService, error) {
	return jwtService{
		jwtKey: jwtKey,
	}, nil
}

// GenerateToken creates a new JWT for a given userID
func (j jwtService) GenerateToken(id int) (string, error) {
	expirationTime := time.Now().Add(consts.JWTExpiryDuration)

	claims := &types.Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.jwtKey)
	return tokenString, err
}
