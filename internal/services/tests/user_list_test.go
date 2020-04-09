package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

var userListTests = []*userTestCase{
	&userTestCase{
		name: "begin fail",
		want: userResult{err: errSome},
	},
	&userTestCase{
		name:    "repository fail",
		repoRet: userResult{err: errSome},
		want:    userResult{err: errSome},
	},
	&userTestCase{
		name:    "success",
		repoRet: userResult{err: nil, users: validUsers},
		want:    userResult{err: nil, users: validUsers},
	},
}

func TestList(t *testing.T) {
	for _, test := range userListTests {
		t.Run(test.name, func(t *testing.T) {
			singleTestList(t, test)
		})
	}
}

// singleTestList performs single test
func singleTestList(t *testing.T, test *userTestCase) {
	// You need a database lock to perform some database actions of service
	mockDB, mock := iniSqlMock(t)
	defer mockDB.Close()
	sqlExpectations(mock, test)

	//You need a repository mock to process the service calls
	mockController, repository, service := initService(t)
	defer mockController.Finish()
	repoListExpectations(repository, test)

	users, err := service.List()
	assert.Equal(t, test.want.err, err)
	assert.Equal(t, test.want.users, users)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func repoListExpectations(repository *mocks.MockUser, test *userTestCase) {
	if test.name == "begin fail" {
		return
	}
	repository.EXPECT().
		List(gomock.Not(gomock.Nil())).
		Return(test.repoRet.users, test.repoRet.err)
}
