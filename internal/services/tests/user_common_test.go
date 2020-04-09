package services_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	"github.com/goofinator/usersHttp/internal/services"
)

type userResult struct {
	err error
}

type userTestCase struct {
	name    string
	user    *model.User
	id      int
	repoRet userResult
	want    userResult
}

var validUser = &model.User{
	ID:        1,
	Age:       2,
	Name:      "Vasya",
	Lastname:  "Pupkin",
	Birthdate: time.Now(),
}

func initService(t *testing.T) (*gomock.Controller, *mocks.MockUser, services.User) {
	mockController := gomock.NewController(t)
	repository := mocks.NewMockUser(mockController)
	service := services.NewUser(repository)
	return mockController, repository, service
}
