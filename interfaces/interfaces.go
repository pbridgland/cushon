package interfaces

import (
	"cushon/types"
)

// DataRepo is used to get data from the DB
type DataRepo interface {
	Funds() (types.Funds, error)
	UsersWithUsername(string) (types.Users, error)
	MakeInvestment(userID int, fundID int, amount int) error
}

// JWTService is used to generate JWTs
type JWTService interface {
	GenerateToken(id int) (string, error)
}

// LoginService is used to log in a user
type LoginService interface {
	Login(credentials types.Credentials) (int, error)
}
