package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goofinator/usersHttp/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

var userReplaceTests = []*userTestCase{
	&userTestCase{
		name: "begin fail",
		id:   1,
		user: validUser,
		want: userResult{err: errSome},
	},
	&userTestCase{
		name:    "repository fail",
		id:      1,
		user:    validUser,
		repoRet: userResult{err: errSome},
		want:    userResult{err: errSome},
	},
	&userTestCase{
		name:    "success",
		id:      1,
		user:    validUser,
		repoRet: userResult{err: nil},
		want:    userResult{err: nil},
	},
}

func TestReplace(t *testing.T) {
	for _, test := range userReplaceTests {
		t.Run(test.name, func(t *testing.T) {
			singleTestReplace(t, test)
		})
	}
}

// singleTestReplace performs single test
func singleTestReplace(t *testing.T, test *userTestCase) {
	// You need a database lock to perform some database actions of service
	mockDB, mock := iniSqlMock(t)
	defer mockDB.Close()
	sqlExpectations(mock, test)

	//You need a repository mock to process the service calls
	mockController, repository, service := initService(t)
	defer mockController.Finish()
	repoReplaceExpectations(repository, test)

	err := service.Replace(test.id, test.user)
	assert.Equal(t, test.want.err, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func repoReplaceExpectations(repository *mocks.MockUser, test *userTestCase) {
	if test.name == "begin fail" {
		return
	}
	repository.EXPECT().
		Replace(gomock.Not(gomock.Nil()), test.id, test.user).
		Return(test.repoRet.err)
}
