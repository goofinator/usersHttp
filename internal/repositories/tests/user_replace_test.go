package repositories_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/goofinator/usersHttp/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var userReplaceTests = []*userTestCase{
	&userTestCase{
		name: "update fail",
		id:   1,
		user: validUser,
		txRet: userResult{err: errSome,
			result: sqlmock.NewResult(0, 1)},
		want: userResult{err: errSome},
	},
	&userTestCase{
		name: "0 rows",
		id:   1,
		user: validUser,
		txRet: userResult{
			err:    nil,
			result: sqlmock.NewResult(0, 0),
		},
		want: userResult{err: err0LineResult},
	},
	&userTestCase{
		name: "2 rows",
		id:   1,
		user: validUser,
		txRet: userResult{
			err:    nil,
			result: sqlmock.NewResult(0, 2),
		},
		want: userResult{err: err2LineResult},
	},
	&userTestCase{
		name: "success",
		id:   1,
		user: validUser,
		txRet: userResult{
			err:    nil,
			result: sqlmock.NewResult(0, 1),
		},
		want: userResult{err: nil},
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
	sqlReplaceExpectations(mock, test)

	repository := repositories.NewUser()
	tx, err := mockDB.Begin()
	require.NoError(t, err)

	err = repository.Replace(tx, test.id, test.user)
	assert.Equal(t, test.want.err, err)

	err = tx.Commit()
	require.NoError(t, err)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}

func sqlReplaceExpectations(mock sqlmock.Sqlmock, test *userTestCase) {
	mock.ExpectBegin()

	reqRegExp := &strings.Builder{}
	fmt.Fprint(reqRegExp, "^UPDATE http_users SET name=\\$1, ")
	fmt.Fprint(reqRegExp, "lastname=\\$2, birthdate=\\$3 WHERE +id=\\$4$")

	ex := mock.ExpectExec(reqRegExp.String()).
		WithArgs(test.user.Name, test.user.Lastname, test.user.Birthdate, test.id)

	if test.name == "update fail" {
		ex.WillReturnError(test.txRet.err)
	} else {
		ex.WillReturnResult(test.txRet.result)
	}

	mock.ExpectCommit()
}
