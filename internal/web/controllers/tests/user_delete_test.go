package controllers_test

import (
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/services/mocks"
	"github.com/goofinator/usersHttp/internal/web/binders"
	"github.com/goofinator/usersHttp/internal/web/controllers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testsDelete = []*commonUserTestCase{
	{
		name:       "wrong URL's id format",
		url:        "/users/A",
		wantStatus: http.StatusNotFound,
		wantBodyRE: "^404 page not found",
		mockRetErr: nil,
	},
	{
		name:       "db error",
		url:        "/users/1",
		wantStatus: http.StatusBadRequest,
		wantBodyRE: "^some error",
		mockRetErr: someError,
	},
	{
		name:       "success",
		url:        "/users/1",
		wantStatus: http.StatusOK,
		wantBodyRE: "",
		mockRetErr: nil,
	},
}

func TestDeleteHandler(t *testing.T) {
	for _, test := range testsDelete {
		t.Run(test.name, func(t *testing.T) {
			//You need a Service mock to process the request
			mockController := gomock.NewController(t)
			defer mockController.Finish()
			service := mocks.NewMockUser(mockController)
			controller := controllers.NewUser(service)

			setDeleteExpectations(service, test)

			req, err := http.NewRequest("DELETE", test.url, nil)
			require.NoError(t, err)

			rr := handleRequest(req, binders.IDBinder(controller.Delete),
				&handlingParams{route: "/users/{id:[0-9]+}", method: "DELETE"})

			assert.Equal(t, test.wantStatus, rr.Code)
			assert.Regexp(t, regexp.MustCompile(test.wantBodyRE), rr.Body.String())
		})
	}
}

func setDeleteExpectations(controller *mocks.MockUser, test *commonUserTestCase) {
	if strings.HasPrefix(test.name, "wrong URL") {
		return
	}

	controller.EXPECT().
		Delete(1).Return(test.mockRetErr)
}
