package services_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/model"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	"github.com/goofinator/usersHttp/internal/services"
)

type userResult struct {
	err   error
	users []*model.User
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

var validUsers = []*model.User{
	&model.User{
		ID:        1,
		Age:       2,
		Name:      "Vasya",
		Lastname:  "Pupkin",
		Birthdate: time.Now(),
	},
	&model.User{
		ID:        3,
		Age:       4,
		Name:      "Venya",
		Lastname:  "Levkin",
		Birthdate: time.Now(),
	},
}

func initService(t *testing.T) (*gomock.Controller, *mocks.MockUser, services.User) {
	mockController := gomock.NewController(t)
	repository := mocks.NewMockUser(mockController)
	service := services.NewUser(repository)
	return mockController, repository, service
}

func sqlExpectations(mock sqlmock.Sqlmock, test *userTestCase) {
	ex := mock.ExpectBegin()
	if test.name == "begin fail" {
		ex.WillReturnError(errSome)
		return
	}
	if test.name == "repository fail" {
		mock.ExpectRollback()
		return
	}
	mock.ExpectCommit()
}
