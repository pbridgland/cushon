package mocks

import (
	"cushon/types"
)

type DataRepo struct {
	NextFundsResult             types.Funds
	NextFundsErr                error
	NextUsersWithUsernameResult types.Users
	NextUsersWithUsernameErr    error
	NextMakeInvestmentErr       error
}

func (d *DataRepo) UsersWithUsername(string) (types.Users, error) {
	return d.NextUsersWithUsernameResult, d.NextUsersWithUsernameErr
}

func (d *DataRepo) MakeInvestment(userID int, param2 int, amount int) error {
	return d.NextMakeInvestmentErr
}

func (d *DataRepo) Funds() (types.Funds, error) {
	return d.NextFundsResult, d.NextFundsErr
}
