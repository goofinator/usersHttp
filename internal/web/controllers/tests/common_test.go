package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

var someError = errors.New("some error")

type handlingParams struct {
	route  string
	method string
}

func handleRequest(req *http.Request, fnc http.HandlerFunc, hp *handlingParams) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(hp.route, fnc).Methods(hp.method)
	router.ServeHTTP(rr, req)
	return rr
}
