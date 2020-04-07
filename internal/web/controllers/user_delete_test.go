package controllers_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	. "github.com/goofinator/usersHttp/internal/web/controllers"
)

var testsDeleteUser = []*commonTestCase{
	{
		name:       "wrong URL's id format 1",
		url:        "/users/:A",
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on IDFromURL",
		mockRetErr: nil,
	},
	{
		name:       "wrong URL's id format 2",
		url:        "/users/1",
		wantStatus: http.StatusUnprocessableEntity,
		wantBodyRE: "^error on IDFromURL",
		mockRetErr: nil,
	},
	{
		name:       "db error",
		url:        "/users/:1",
		wantStatus: http.StatusInternalServerError,
		wantBodyRE: "^error on DeleteUser: some error",
		mockRetErr: someError,
	},
	{
		name:       "success",
		url:        "/users/:1",
		wantStatus: http.StatusOK,
		wantBodyRE: "",
		mockRetErr: nil,
	},
}

func TestDeleteUserHandler(t *testing.T) {
	for _, test := range testsDeleteUser {
		t.Run(test.name, func(t *testing.T) {
			//You need a Storager mock to process the request
			controller := gomock.NewController(t)
			defer controller.Finish()
			db := mocks.NewMockStorager(controller)

			setDeleteUserExpectations(db, test)

			req, err := http.NewRequest("DELETE", test.url, nil)
			if err != nil {
				t.Fatalf("unexpected fail of NewRequest: %s", err)
			}
			rr := handleRequest(req, db, DeleteUserHandler)

			checkStatus(t, test.wantStatus, rr.Code)
			checkBodyByRE(t, test.wantBodyRE, rr.Body.String())
		})
	}
}

func setDeleteUserExpectations(db *mocks.MockStorager, test *commonTestCase) {
	if strings.HasPrefix(test.name, "wrong URL") {
		return
	}

	db.EXPECT().
		DeleteUser(1).Return(test.mockRetErr)
}
