package controllers

import (
	"fmt"
	"net/http"
)

// DeleteUserHandler handles DELETE request on /users/:id endpoint
// to remove user from a DB
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "//TODO: DeleteUserHandler")

}
