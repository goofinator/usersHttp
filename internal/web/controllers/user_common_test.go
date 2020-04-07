package controllers_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/repositories"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
)

var someError = errors.New("some error")

var (
	jsonValidStr = []byte(fmt.Sprintf(`{"Id": 0,
	"Name": "petya",
	"Lastname": "Pupkin",
	"Age": 22,
	"Birthdate": "%v"}`,
		time.Now().UTC().Format(time.RFC3339)))
	jsonInvalidStr = []byte(`{"Id": 0,
	"Name": "petya",
	"Lastname": "Pupkin",
	"Age": 22,
	"Birthdate": "AAA"}`)
)

type commonTestCase struct {
	name         string
	jsonStr      []byte
	url          string
	wantStatus   int
	wantBodyRE   string
	mockRetErr   error
	mockRetUsers []*model.User
}

func handleRequest(req *http.Request, db *mocks.MockStorager,
	fnc func(w http.ResponseWriter, r *http.Request, db repositories.Storager)) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fnc(w, r, db)
	})
	handler.ServeHTTP(rr, req)
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
