package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goofinator/usersHttp/internal/init/startup"
	"github.com/goofinator/usersHttp/internal/repositories"
	"github.com/goofinator/usersHttp/internal/services"
	"github.com/goofinator/usersHttp/internal/web/binders"
	"github.com/goofinator/usersHttp/internal/web/controllers"
	"github.com/gorilla/mux"
)

// Run starts the web service
func Run(iniData *startup.IniData) {
	router := mux.NewRouter()

	handleRoutes(router)
	http.Handle("/", router)

	err := http.ListenAndServe(fmt.Sprintf(":%d", iniData.Port), nil)
	if err != nil {
		log.Fatalf("unexpected error while ListenAndServe: %s", err)
	}
}

func handleRoutes(router *mux.Router) {
	uc := userController()

	router.HandleFunc("/users", uc.Add).Methods("POST")

	router.HandleFunc("/users", uc.List).Methods("GET")

	router.HandleFunc("/users/{id:[0-9]+}",
		binders.IDBinder(uc.Delete)).Methods("DELETE")

	router.HandleFunc("/users/{id:[0-9]+}",
		binders.IDBinder(uc.Replace)).Methods("PUT")

}

func userController() controllers.User {
	repository := repositories.NewUser()
	service := services.NewUser(repository)
	return controllers.NewUser(service)
}
