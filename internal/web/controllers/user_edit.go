package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goofinator/usersHttp/internal/repositories"
	"github.com/goofinator/usersHttp/internal/utils"
	"github.com/goofinator/usersHttp/internal/web/model"
)

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
