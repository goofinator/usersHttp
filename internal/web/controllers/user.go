package controllers

import (
	"net/http"

	"github.com/goofinator/usersHttp/internal/services"
)

// User interface wrapps the user's controller functions
type User interface {
	Add(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
	Replace(http.ResponseWriter, *http.Request)
}

// NewUser produces a user's controller
func NewUser(service services.User) User {
	return &user{service: service}
}

type user struct {
	service services.User
}

func (u *user) Add(w http.ResponseWriter, r *http.Request) {

}

func (u *user) Delete(w http.ResponseWriter, r *http.Request) {

}

func (u *user) List(w http.ResponseWriter, r *http.Request) {

}

func (u *user) Replace(w http.ResponseWriter, r *http.Request) {

}
