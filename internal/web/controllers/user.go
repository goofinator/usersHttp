package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/services"
	"github.com/goofinator/usersHttp/internal/web/binders"
	"github.com/gorilla/context"
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

// Add invokes user's service to store usere
func (u *user) Add(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.service.Add(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// Delete invokes user's service to delete usere
func (u *user) Delete(w http.ResponseWriter, r *http.Request) {
	id := context.Get(r, binders.ID).(int)

	if err := u.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// List invokes user's service to return list of users
func (u *user) List(w http.ResponseWriter, r *http.Request) {
	users, err := u.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := encodeUsers(w, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Replace invokes user's service to replace data of user with specified id
func (u *user) Replace(w http.ResponseWriter, r *http.Request) {
	id := context.Get(r, binders.ID).(int)
	user, err := decodeUser(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.service.Replace(id, user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

// DecodeUser decodes user from incoming json
func decodeUser(r io.Reader) (*model.User, error) {
	decoder := json.NewDecoder(r)
	var user model.User

	if err := decoder.Decode(&user); err != nil {
		return nil, fmt.Errorf("error on json.Decode: %s", err)
	}
	return &user, nil
}

// EncodeUsers encode users to json
func encodeUsers(w io.Writer, users []*model.User) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(users); err != nil {
		return fmt.Errorf("error on jsonEncode: %s", err)
	}
	return nil
}
