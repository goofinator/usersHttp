package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goofinator/usersHttp/internal/repositories"
	"github.com/goofinator/usersHttp/internal/utils"
	"github.com/goofinator/usersHttp/internal/web/model"
)

// AddUserHandler handles POST request on /users endpoint
// to add user to a DB
func AddUserHandler(w http.ResponseWriter, r *http.Request, db repositories.Storager) {
	decoder := json.NewDecoder(r.Body)
	var user model.User

	if err := decoder.Decode(&user); err != nil {
		msg := fmt.Sprintf("error on json.Decode: %s.", err)
		http.Error(w, msg, http.StatusUnprocessableEntity)
		return
	}

	if err := db.AddUser(&user); err != nil {
		msg := fmt.Sprintf("error on AddUser: %s.", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

// DeleteUserHandler handles DELETE request on /users/:id endpoint
// to remove user from a DB
func DeleteUserHandler(w http.ResponseWriter, r *http.Request, db repositories.Storager) {
	id, err := utils.IDFromURL(r.URL)
	if err != nil {
		msg := fmt.Sprintf("error on IDFromURL: %s.", err)
		http.Error(w, msg, http.StatusUnprocessableEntity)
		return
	}

	if err := db.DeleteUser(id); err != nil {
		msg := fmt.Sprintf("error on DeleteUser: %s.", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

}

// EditUserHandler handles PUT request on /users/:id endpoint
// to edit user in a DB
func EditUserHandler(w http.ResponseWriter, r *http.Request, db repositories.Storager) {
	id, err := utils.IDFromURL(r.URL)
	if err != nil {
		msg := fmt.Sprintf("error on IDFromURL: %s.", err)
		http.Error(w, msg, http.StatusUnprocessableEntity)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		msg := fmt.Sprintf("error on json.Decode: %s.", err)
		http.Error(w, msg, http.StatusUnprocessableEntity)
		return
	}

	if err := db.EditUser(id, &user); err != nil {
		msg := fmt.Sprintf("error on EditUser: %s.", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

// ListUsersHandler handles GET request on /users endpoint
// to get all users from a DB
func ListUsersHandler(w http.ResponseWriter, r *http.Request, db repositories.Storager) {
	w.Header().Set("Content-Type", "application/json")
	users, err := db.GetUsers()
	if err != nil {
		msg := fmt.Sprintf("error on GetUsers: %s.", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(users)
	if err != nil {
		msg := fmt.Sprintf("error on json.MarshalIndent: %s.", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}
