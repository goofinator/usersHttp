package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/services/mocks"
	"github.com/goofinator/usersHttp/internal/web/binders"
	"github.com/goofinator/usersHttp/internal/web/controllers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testsReplace = []*commonUserTestCase{
	{
		name:       "wrong URL's id format",
		url:        "/users/A",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusNotFound,
		wantBodyRE: "^404 page not found",
		mockRetErr: nil,
	},
	{
		name:       "broken json",
		url:        "/users/1",
		jsonStr:    jsonInvalidStr,
		wantStatus: http.StatusBadRequest,
		wantBodyRE: "^error on json\\.Decode: parsing time",
		mockRetErr: nil,
	},
	{
		name:       "db error",
		url:        "/users/1",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusBadRequest,
		wantBodyRE: "some error",
		mockRetErr: someError,
	},
	{
		name:       "success",
		url:        "/users/1",
		jsonStr:    jsonValidStr,
		wantStatus: http.StatusOK,
		wantBodyRE: "",
		mockRetErr: nil,
	},
}

func TestReplaceHandler(t *testing.T) {
	for _, test := range testsReplace {
		t.Run(test.name, func(t *testing.T) {
			//You need a Service mock to process the request
			mockController := gomock.NewController(t)
			defer mockController.Finish()
			service := mocks.NewMockUser(mockController)
			controller := controllers.NewUser(service)

			setReplaceExpectations(t, service, test)

			req, err := http.NewRequest("PUT", test.url, bytes.NewBuffer(test.jsonStr))
			require.NoError(t, err)

			rr := handleRequest(req, binders.IDBinder(controller.Replace),
				&handlingParams{route: "/users/{id:[0-9]+}", method: "PUT"})

			assert.Equal(t, test.wantStatus, rr.Code)
			assert.Regexp(t, regexp.MustCompile(test.wantBodyRE), rr.Body.String())
		})
	}
}

func setReplaceExpectations(t *testing.T, controller *mocks.MockUser, test *commonUserTestCase) {
	if test.name == "broken json" || strings.HasPrefix(test.name, "wrong URL") {
		return
	}

	var user model.User
	if err := json.Unmarshal(jsonValidStr, &user); err != nil {
		t.Fatalf("unexpected fail of Unmarshal: %s", err)
	}

	controller.EXPECT().
		Replace(1, &user).Return(test.mockRetErr)
}
