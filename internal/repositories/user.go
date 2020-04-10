//go:generate mockgen -destination=./mocks/mock_user.go -package=mocks github.com/goofinator/usersHttp/internal/repositories User
//go:generate goimports -w ./mocks/mock_user.go

package repositories

import (
	"database/sql"

	"github.com/goofinator/usersHttp/internal/datasource"
	"github.com/goofinator/usersHttp/internal/model"
)

// User interface wraps user's repository functions
// to perform simple operation with given transaction tx
type User interface {
	Add(tx *sql.Tx, user *model.User) error
	Delete(tx *sql.Tx, id int) error
	List(tx *sql.Tx) ([]*model.User, error)
	Replace(tx *sql.Tx, id int, user *model.User) error
}

// NewUser creates a new User repository
func NewUser() User {
	return &user{}
}

type user struct{}

// Add inserts User in repository
func (u *user) Add(tx *sql.Tx, user *model.User) error {
	query := `INSERT INTO http_users 	(id, name, lastname, birthdate)
	VALUES(DEFAULT,$1,$2,$3)`
	result, err := tx.Exec(query, user.Name, user.Lastname, user.Birthdate)

	if err := datasource.CheckResult(1, result, err); err != nil {
		return err
	}
	return nil
}

// Delete removes user from repository
func (u *user) Delete(tx *sql.Tx, id int) error {
	query := "DELETE FROM http_users WHERE id=$1"
	result, err := tx.Exec(query, id)

	if err := datasource.CheckResult(1, result, err); err != nil {
		return err
	}
	return nil
}

// List returns all users from repository
func (u *user) List(tx *sql.Tx) ([]*model.User, error) {
	query := `SELECT id, name, lastname, birthdate, EXTRACT(YEAR FROM AGE(birthdate)) 
	AS age FROM http_users ORDER BY ID`
	rows, err := tx.Query(query)

	users, err := processRows(rows, err)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Replace put user's data in repository to Row with specified id
func (u *user) Replace(tx *sql.Tx, id int, user *model.User) error {
	query := `UPDATE http_users SET name=$1, lastname=$2, birthdate=$3 WHERE id=$4`
	result, err := tx.Exec(query, user.Name, user.Lastname, user.Birthdate, id)

	if err := datasource.CheckResult(1, result, err); err != nil {
		return err
	}
	return nil
}

func processRows(rows *sql.Rows, err error) ([]*model.User, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*model.User, 0)

	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Birthdate, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
