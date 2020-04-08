package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/services/mocks"
	"github.com/goofinator/usersHttp/internal/web/controllers"
)

var testsAdd = []*commonUserTestCase{
	{
		name:       "broken json",
		jsonStr:    jsonInvalidStr,
		wantStatus: http.StatusBadRequest,
		wantBodyRE: "^error on json\\.Decode: parsing time",
		mockRetErr: nil,
	},
	{
		name:       "db error",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusBadRequest,
		wantBodyRE: "^some error",
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

func TestAddHandler(t *testing.T) {
	for _, test := range testsAdd {
		t.Run(test.name, func(t *testing.T) {
			//You need a Service mock to process the request
			mockController := gomock.NewController(t)
			defer mockController.Finish()
			service := mocks.NewMockUser(mockController)
			controller := controllers.NewUser(service)

			setAddExpectations(t, service, test)

			req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(test.jsonStr))
			if err != nil {
				t.Fatalf("unexpected fail of NewRequest: %s", err)
			}
			req.Header.Set("Content-Type", "application/json")
			rr := handleRequest(req, controller.Add,
				&handlingParams{route: "/users", method: "POST"})

			checkStatus(t, test.wantStatus, rr.Code)
			checkBodyByRE(t, test.wantBodyRE, rr.Body.String())
		})
	}
}

func setAddExpectations(t *testing.T, controller *mocks.MockUser, test *commonUserTestCase) {
	if test.name == "broken json" {
		return
	}

	var user model.User
	if err := json.Unmarshal(jsonValidStr, &user); err != nil {
		t.Fatalf("unexpected fail of Unmarshal: %s", err)
	}

	controller.EXPECT().
		Add(&user).Return(test.mockRetErr)
}
