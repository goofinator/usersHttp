package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	. "github.com/goofinator/usersHttp/internal/web/controllers"
	"github.com/goofinator/usersHttp/internal/web/model"
)

var testsEditUser = []*commonTestCase{
	{
		name:       "wrong URL's id format 1",
		url:        "/users/:A",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on IDFromURL",
		mockRet:    nil,
	},
	{
		name:       "wrong URL's id format 2",
		url:        "/users/1",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on IDFromURL",
		mockRet:    nil,
	},
	{
		name:       "broken json",
		url:        "/users/:1",
		jsonStr:    jsonInvalidStr,
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on json\\.Decode: parsing time",
		mockRet:    nil,
	},
	{
		name:       "db error",
		url:        "/users/:1",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusInternalServerError,
		wantBodyRE: "error on EditUser: some error",
		mockRet:    someError,
	},
	{
		name:       "success",
		url:        "/users/:1",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusOK,
		wantBodyRE: "",
		mockRet:    nil,
	},
}

func TestEditUserHandler(t *testing.T) {
	for _, test := range testsEditUser {
		t.Run(test.name, func(t *testing.T) {
			//You need a Storager mock to process the request
			controller := gomock.NewController(t)
			defer controller.Finish()
			db := mocks.NewMockStorager(controller)

			setEditUserExpectations(db, test)

			req, err := http.NewRequest("PUT", test.url, bytes.NewBuffer(test.jsonStr))
			if err != nil {
				t.Fatalf("unexpected fail of NewRequest: %s", err)
			}
			rr := handleRequest(req, db, EditUserHandler)

			checkStatus(t, test.wantStatus, rr.Code)

			checkBodyByRE(t, test.wantBodyRE, rr.Body.String())
		})
	}
}

func setEditUserExpectations(db *mocks.MockStorager, test *commonTestCase) {
	if test.name == "broken json" || strings.HasPrefix(test.name, "wrong URL") {
		return
	}

	var user model.User
	json.Unmarshal(jsonValidStr, &user)

	db.EXPECT().
		EditUser(1, &user).Return(test.mockRet)
}
