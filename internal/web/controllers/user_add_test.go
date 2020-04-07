package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	. "github.com/goofinator/usersHttp/internal/web/controllers"
)

var testsAddUser = []*commonTestCase{
	{
		name:       "broken json",
		jsonStr:    jsonInvalidStr,
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on json\\.Decode: parsing time",
		mockRetErr: nil,
	},
	{
		name:       "db error",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusInternalServerError,
		wantBodyRE: "^error on AddUser: some error",
		mockRetErr: someError,
	},
	{
		name:       "success",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusOK,
		wantBodyRE: "^$",
		mockRetErr: nil,
	},
}

func TestAddUserHandler(t *testing.T) {
	for _, test := range testsAddUser {
		t.Run(test.name, func(t *testing.T) {
			//You need a Storager mock to process the request
			controller := gomock.NewController(t)
			defer controller.Finish()
			db := mocks.NewMockStorager(controller)

			setAddUserExpectations(t, db, test)

			req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(test.jsonStr))
			if err != nil {
				t.Fatalf("unexpected fail of NewRequest: %s", err)
			}
			req.Header.Set("Content-Type", "application/json")
			rr := handleRequest(req, db, AddUserHandler)

			checkStatus(t, test.wantStatus, rr.Code)
			checkBodyByRE(t, test.wantBodyRE, rr.Body.String())
		})
	}
}

func setAddUserExpectations(t *testing.T, db *mocks.MockStorager, test *commonTestCase) {
	if test.name == "broken json" {
		return
	}

	var user model.User
	if err := json.Unmarshal(jsonValidStr, &user); err != nil {
		t.Fatalf("unexpected fail of Unmarshal: %s", err)
	}

	db.EXPECT().
		AddUser(&user).Return(test.mockRetErr)
}
