package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

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

func checkStatus(t *testing.T, want, got int) {
	if want != got {
		t.Errorf("unexpected status code:\nwant: %v\ngot: %v", want, got)
	}
}

func checkBodyByRE(t *testing.T, wantRE, got string) {
	matched, err := regexp.MatchString(wantRE, got)
	if err != nil {
		t.Fatalf("unexpected fail on MatchString: %v", err)
	}
	if !matched {
		t.Errorf("unexpected response body:\nwant: match to regexp: %q\ngot: %q", wantRE, got)
	}
}
