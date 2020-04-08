//go:generate mockgen -destination=./mocks/mock_user.go -package=mocks github.com/goofinator/usersHttp/internal/services User
//go:generate goimports -w ./mocks/mock_user.go

package services

import (
	"github.com/goofinator/usersHttp/internal/datasource"
	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/repositories"
)

// User wraps the user's service function
type User interface {
	Add(user *model.User) error
	Delete(id int) error
	List() ([]*model.User, error)
	Replace(id int, user *model.User) error
}

// NewUser creates User service
func NewUser(repository repositories.User) User {
	return &user{repository: repository}
}

type user struct {
	repository repositories.User
}

func (u *user) Add(user *model.User) (err error) {
	tx, err := datasource.SQL.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = datasource.CloseTransaction(tx, err)
	}()

	if err := u.repository.Add(tx, user); err != nil {
		return err
	}
	return nil
}

func (u *user) Delete(id int) (err error) {
	tx, err := datasource.SQL.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = datasource.CloseTransaction(tx, err)
	}()

	if err := u.repository.Delete(tx, id); err != nil {
		return err
	}
	return nil
}

func (u *user) List() (users []*model.User, err error) {
	tx, err := datasource.SQL.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = datasource.CloseTransaction(tx, err)
	}()

	users, err = u.repository.List(tx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *user) Replace(id int, user *model.User) (err error) {
	tx, err := datasource.SQL.Begin()
	if err != nil {
		return err
	}
	defer func() {
		err = datasource.CloseTransaction(tx, err)
	}()

	if err := u.repository.Replace(tx, id, user); err != nil {
		return err
	}
	return nil
}
