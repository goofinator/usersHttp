package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	. "github.com/goofinator/usersHttp/internal/web/controllers"
	"github.com/goofinator/usersHttp/internal/web/model"
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
	name       string
	jsonStr    []byte
	wantStatus int
	wantBodyRE string
	mockRet    error
}

var tests = []*commonTestCase{
	{
		name:       "broken json",
		jsonStr:    jsonInvalidStr,
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on json\\.Decode: parsing time",
		mockRet:    nil,
	},
	{
		name:       "db error",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusInternalServerError,
		wantBodyRE: "",
		mockRet:    someError,
	},
	{
		name:       "success",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusOK,
		wantBodyRE: "",
		mockRet:    nil,
	},
}

func TestAddUserHandler(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//You need a Storager mock to process the request
			controller := gomock.NewController(t)
			defer controller.Finish()
			db := mocks.NewMockStorager(controller)

			if bytes.Equal(test.jsonStr, jsonValidStr) {
				setAddUserExpectations(db, test)
			}

			req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(test.jsonStr))
			if err != nil {
				t.Fatalf("unexpected fail of NewRequest: %s", err)
			}
			req.Header.Set("Content-Type", "application/json")
			rr := handleRequest(req, db)

			checkStatus(t, test.wantStatus, rr.Code)

			checkBodyByRE(t, test.wantBodyRE, rr.Body.String())
		})
	}
}

func setAddUserExpectations(db *mocks.MockStorager, test *commonTestCase) {
	var user model.User
	json.Unmarshal(jsonValidStr, &user)

	db.EXPECT().
		AddUser(&user).Return(test.mockRet)
}

func handleRequest(req *http.Request, db *mocks.MockStorager) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddUserHandler(w, r, db)
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
