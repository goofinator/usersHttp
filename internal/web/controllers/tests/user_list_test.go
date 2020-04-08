package controllers_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/services/mocks"
	"github.com/goofinator/usersHttp/internal/web/controllers"
)

var testsList = []*commonUserTestCase{
	{
		name:         "db error",
		wantStatus:   http.StatusBadRequest,
		wantBodyRE:   "^some error",
		mockRetErr:   someError,
		mockRetUsers: nil,
	},
	{
		name:       "success",
		wantStatus: http.StatusOK,
		wantBodyRE: ".",
		mockRetErr: nil,
		mockRetUsers: []*model.User{
			&model.User{
				ID:        1,
				Name:      "Cheech",
				Lastname:  "Marin",
				Age:       73,
				Birthdate: time.Now(),
			},
			&model.User{
				ID:        2,
				Name:      "Kevin",
				Lastname:  "Costner",
				Age:       65,
				Birthdate: time.Now(),
			},
		},
	},
}

func TestListHandler(t *testing.T) {
	for _, test := range testsList {
		t.Run(test.name, func(t *testing.T) {
			//You need a Service mock to process the request
			mockController := gomock.NewController(t)
			defer mockController.Finish()
			service := mocks.NewMockUser(mockController)
			controller := controllers.NewUser(service)

			setListExpectations(service, test)

			req, err := http.NewRequest("GET", "/users", nil)
			if err != nil {
				t.Fatalf("unexpected fail of NewRequest: %s", err)
			}
			rr := handleRequest(req, controller.List,
				&handlingParams{route: "/users", method: "GET"})

			checkStatus(t, test.wantStatus, rr.Code)
			checkBodyByRE(t, test.wantBodyRE, rr.Body.String())
			if rr.Code == http.StatusOK {
				checkUsers(t, test, rr.Body.Bytes())
			}
		})
	}
}

func setListExpectations(controller *mocks.MockUser, test *commonUserTestCase) {
	controller.EXPECT().
		List().Return(test.mockRetUsers, test.mockRetErr)
}

func checkUsers(t *testing.T, test *commonUserTestCase, body []byte) {
	wantJson, err := json.MarshalIndent(test.mockRetUsers, "", "    ")
	if err != nil {
		t.Fatalf("unexpected fail of MarshalIndent: %s", err)
	}
	sWantJson := strings.TrimRight(string(wantJson), "\n ")
	sBody := strings.TrimRight(string(body), "\n ")

	if sWantJson != sBody {
		t.Errorf("unexpected response body:\nwant: %q\ngot: %q", sWantJson, sBody)
	}
}
