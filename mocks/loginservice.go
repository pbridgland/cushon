package mocks

import "cushon/types"

type LoginService struct {
	NextLoginResult int
	NextLoginErr    error
}

// Login implements interfaces.LoginService.
func (l *LoginService) Login(credentials types.Credentials) (int, error) {
	return l.NextLoginResult, l.NextLoginErr
}
