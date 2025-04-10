package repos

import (
	"cushon/types"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "cushon"
)

type data struct {
	db *sql.DB
}

// NewDataRepo creates a new instance of the Data Repository
func NewDataRepo() (data, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return data{}, err
	}

	err = db.Ping()
	if err != nil {
		return data{}, err
	}
	return data{db: db}, nil
}

// Funds returns a slice of all funds in the DB
func (d *data) Funds() (types.Funds, error) {
	rows, err := d.db.Query("SELECT * FROM get_all_funds()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	funds := make(types.Funds, 0)
	for rows.Next() {
		fund := types.Fund{}
		err = rows.Scan(&fund.ID, &fund.Name)
		if err != nil {
			return nil, err
		}
		funds = append(funds, fund)
	}

	return funds, nil
}

// UsersWithUsername returns a slice of all users in the DB that have the given username
func (d *data) UsersWithUsername(username string) (types.Users, error) {
	query := fmt.Sprintf("SELECT * FROM get_user_by_username('%s')", username)

	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make(types.Users, 0)
	for rows.Next() {
		var user types.User
		err = rows.Scan(&user.ID, &user.Username, &user.HashedPassword)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// MakeInvestment records the given user making an investment in the given fund for the given amount
func (d *data) MakeInvestment(userID int, fundID int, amount int) error {
	query := fmt.Sprintf("CALL add_investment(%d,%d,%d);", userID, fundID, amount)

	rows, err := d.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
