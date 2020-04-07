package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	. "github.com/goofinator/usersHttp/internal/web/controllers"
)

var testsEditUser = []*commonTestCase{
	{
		name:       "wrong URL's id format 1",
		url:        "/users/:A",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on IDFromURL",
		mockRetErr: nil,
	},
	{
		name:       "wrong URL's id format 2",
		url:        "/users/1",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on IDFromURL",
		mockRetErr: nil,
	},
	{
		name:       "broken json",
		url:        "/users/:1",
		jsonStr:    jsonInvalidStr,
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on json\\.Decode: parsing time",
		mockRetErr: nil,
	},
	{
		name:       "db error",
		url:        "/users/:1",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusInternalServerError,
		wantBodyRE: "error on EditUser: some error",
		mockRetErr: someError,
	},
	{
		name:       "success",
		url:        "/users/:1",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusOK,
		wantBodyRE: "",
		mockRetErr: nil,
	},
}

func TestEditUserHandler(t *testing.T) {
	for _, test := range testsEditUser {
		t.Run(test.name, func(t *testing.T) {
			//You need a Storager mock to process the request
			controller := gomock.NewController(t)
			defer controller.Finish()
			db := mocks.NewMockStorager(controller)

			setEditUserExpectations(t, db, test)

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

func setEditUserExpectations(t *testing.T, db *mocks.MockStorager, test *commonTestCase) {
	if test.name == "broken json" || strings.HasPrefix(test.name, "wrong URL") {
		return
	}

	var user model.User
	if err := json.Unmarshal(jsonValidStr, &user); err != nil {
		t.Fatalf("unexpected fail of Unmarshal: %s", err)
	}

	db.EXPECT().
		EditUser(1, &user).Return(test.mockRetErr)
}
