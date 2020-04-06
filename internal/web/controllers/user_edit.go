package controllers

import (
	"fmt"
	"net/http"
)

// EditUserHandler handles PUT request on /users/:id endpoint
// to edit user in a DB
func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "//TODO: EditUserHandler")
}
