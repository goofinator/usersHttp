package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goofinator/usersHttp/internal/init/startup"
	"github.com/goofinator/usersHttp/internal/web/controllers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Run starts the web service
func Run(iniData *startup.IniData) {
	router := mux.NewRouter()

	db := startup.InitDB(iniData)

	handleRoutes(router, db)
	http.Handle("/", router)

	err := http.ListenAndServe(fmt.Sprintf(":%d", iniData.Port), nil)
	if err != nil {
		log.Fatalf("unexpected error while ListenAndServe: %s", err)
	}
}

func handleRoutes(router *mux.Router, db *sqlx.DB) {
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.AddUserHandler(w, r)
	}).Methods("POST")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.ListUsersHandler(w, r)
	}).Methods("GET")

	router.HandleFunc("/users/{id::[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUserHandler(w, r)
	}).Methods("DELETE")

	router.HandleFunc("/users/{id::[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		controllers.EditUserHandler(w, r)
	}).Methods("PUT")
}
