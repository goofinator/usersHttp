package services

import (
	"fmt"

	"github.com/goofinator/usersHttp/internal/model"
)

// User wraps the user's service function
type User interface {
	Add(user *model.User) error
	Delete(id int) error
	List() ([]*model.User, error)
	Replace(id int, user *model.User) error
}

// NewUser creates User service
func NewUser() User {
	return &user{}
}

type user struct{}

func (u *user) Add(user *model.User) error {
	return fmt.Errorf("not implemented")
}

func (u *user) Delete(id int) error {
	return fmt.Errorf("not implemented")
}

func (u *user) List() ([]*model.User, error) {
	return nil, fmt.Errorf("not implemented")
}

func (u *user) Replace(id int, user *model.User) error {
	return fmt.Errorf("not implemented")
}
