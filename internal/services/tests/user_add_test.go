package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

var userAddTests = []*userTestCase{
	&userTestCase{
		name: "begin fail",
		user: validUser,
		want: userResult{err: errSome},
	},
	&userTestCase{
		name:    "repository fail",
		user:    validUser,
		repoRet: userResult{err: errSome},
		want:    userResult{err: errSome},
	},
	&userTestCase{
		name:    "success",
		user:    validUser,
		repoRet: userResult{err: nil},
		want:    userResult{err: nil},
	},
}

func TestAdd(t *testing.T) {
	for _, test := range userAddTests {
		t.Run(test.name, func(t *testing.T) {
			singleTestAdd(t, test)
		})
	}
}

// singleTestAdd performs single test
func singleTestAdd(t *testing.T, test *userTestCase) {
	// You need a database lock to perform some database actions of service
	mockDB, mock := iniSqlMock(t)
	defer mockDB.Close()
	sqlExpectations(mock, test)

	//You need a repository mock to process the service calls
	mockController, repository, service := initService(t)
	defer mockController.Finish()
	repoAddExpectations(repository, test)

	err := service.Add(test.user)
	assert.Equal(t, test.want.err, err)
}

func repoAddExpectations(repository *mocks.MockUser, test *userTestCase) {
	if test.name == "begin fail" {
		return
	}
	repository.EXPECT().
		Add(gomock.Not(gomock.Nil()), test.user).
		Return(test.repoRet.err)
}
