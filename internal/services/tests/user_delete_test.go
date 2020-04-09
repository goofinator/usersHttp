package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

var userDeleteTests = []*userTestCase{
	&userTestCase{
		name: "begin fail",
		id:   1,
		want: userResult{err: errSome},
	},
	&userTestCase{
		name:    "repository fail",
		id:      1,
		repoRet: userResult{err: errSome},
		want:    userResult{err: errSome},
	},
	&userTestCase{
		name:    "success",
		id:      1,
		repoRet: userResult{err: nil},
		want:    userResult{err: nil},
	},
}

func TestDelete(t *testing.T) {
	for _, test := range userDeleteTests {
		t.Run(test.name, func(t *testing.T) {
			singleTestDelete(t, test)
		})
	}
}

// singleTestDelete performs single test
func singleTestDelete(t *testing.T, test *userTestCase) {
	// You need a database lock to perform some database actions of service
	mockDB, mock := iniSqlMock(t)
	defer mockDB.Close()
	sqlExpectations(mock, test)

	//You need a repository mock to process the service calls
	mockController, repository, service := initService(t)
	defer mockController.Finish()
	repoDeleteExpectations(repository, test)

	err := service.Delete(test.id)
	assert.Equal(t, test.want.err, err)
}

func repoDeleteExpectations(repository *mocks.MockUser, test *userTestCase) {
	if test.name == "begin fail" {
		return
	}
	repository.EXPECT().
		Delete(gomock.Not(gomock.Nil()), test.id).
		Return(test.repoRet.err)
}
