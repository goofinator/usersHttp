package repositories

import (
	"database/sql"
	"fmt"

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

func (u *user) Add(tx *sql.Tx, user *model.User) error {
	fmt.Println("Add: ", user)
	return fmt.Errorf("not implemented")
}

func (u *user) Delete(tx *sql.Tx, id int) error {
	fmt.Println("Delete: ", id)
	return fmt.Errorf("not implemented")
}
func (u *user) List(tx *sql.Tx) ([]*model.User, error) {
	fmt.Println("List")
	return nil, fmt.Errorf("not implemented")
}
func (u *user) Replace(tx *sql.Tx, id int, user *model.User) error {
	fmt.Println("Replace: ", id, user)
	return fmt.Errorf("not implemented")
}
