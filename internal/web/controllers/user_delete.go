package controllers

import (
	"fmt"
	"net/http"

	"github.com/goofinator/usersHttp/internal/repositories"
	"github.com/goofinator/usersHttp/internal/utils"
)

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
