package controllers

import (
	"fmt"
	"net/http"
)

// ListUsersHandler handles GET request on /users endpoint
// to get all users from a DB
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "//TODO: ListUsersHandler")
}
