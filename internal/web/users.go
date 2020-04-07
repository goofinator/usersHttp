package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goofinator/usersHttp/internal/init/startup"
	"github.com/goofinator/usersHttp/internal/repositories"
	"github.com/goofinator/usersHttp/internal/web/controllers"
	"github.com/gorilla/mux"
)

// Run starts the web service
func Run(iniData *startup.IniData) {
	router := mux.NewRouter()

	db := repositories.New(iniData)
	defer db.Close()

	handleRoutes(router, db)
	http.Handle("/", router)

	err := http.ListenAndServe(fmt.Sprintf(":%d", iniData.Port), nil)
	if err != nil {
		log.Fatalf("unexpected error while ListenAndServe: %s", err)
	}
}

func handleRoutes(router *mux.Router, db repositories.Storager) {
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.AddUserHandler(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		controllers.ListUsersHandler(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/users/{id::[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteUserHandler(w, r, db)
	}).Methods("DELETE")

	router.HandleFunc("/users/{id::[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		controllers.EditUserHandler(w, r, db)
	}).Methods("PUT")
}
