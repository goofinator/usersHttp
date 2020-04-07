package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goofinator/usersHttp/internal/repositories"
)

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

	prettyJSON, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		msg := fmt.Sprintf("error on json.MarshalIndent: %s.", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(prettyJSON))

}
