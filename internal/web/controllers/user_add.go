package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goofinator/usersHttp/internal/repositories"
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
