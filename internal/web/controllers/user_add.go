package controllers

import (
	"fmt"
	"net/http"
)

// AddUserHandler handles POST request on /users endpoint
// to add user to a DB
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "//TODO: AddUserHandler")
}
