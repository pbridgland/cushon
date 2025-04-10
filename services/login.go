package services

import (
	"cushon/interfaces"
	"cushon/types"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var errMultipleUsersWithSameUsername = errors.New("multiple users found with same username")

type loginService struct {
	dataRepo interfaces.DataRepo
}

// NewLoginService creates a new instance of the Login Service
func NewLoginService(dataRepo interfaces.DataRepo) (loginService, error) {
	return loginService{
		dataRepo: dataRepo,
	}, nil
}

// Login attempts to login a user with the given credentials
// on success it returns the logged in users ID
// on failure it returns an error
func (j loginService) Login(credentials types.Credentials) (int, error) {
	users, err := j.dataRepo.UsersWithUsername(credentials.Username)
	if err != nil {
		return 0, err
	}

	if len(users) != 1 {
		return 0, fmt.Errorf("%w for username %s", errMultipleUsersWithSameUsername, credentials.Username)
	}
	user := users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(credentials.Password))
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
