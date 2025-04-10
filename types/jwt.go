package types

import "github.com/golang-jwt/jwt/v4"

//Claims is the info about a user stored in the JWT, in this case just their user ID
type Claims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
