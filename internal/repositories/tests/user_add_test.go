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

var userAddTests = []*userTestCase{
	&userTestCase{
		name: "insert fail",
		user: validUser,
		txRet: userResult{err: errSome,
			result: sqlmock.NewResult(0, 1)},
		want: userResult{err: errSome},
	},
	&userTestCase{
		name: "0 rows",
		user: validUser,
		txRet: userResult{
			err:    nil,
			result: sqlmock.NewResult(0, 0),
		},
		want: userResult{err: err0LineResult},
	},
	&userTestCase{
		name: "2 rows",
		user: validUser,
		txRet: userResult{
			err:    nil,
			result: sqlmock.NewResult(0, 2),
		},
		want: userResult{err: err2LineResult},
	},
	&userTestCase{
		name: "success",
		user: validUser,
		txRet: userResult{
			err:    nil,
			result: sqlmock.NewResult(0, 1),
		},
		want: userResult{err: nil},
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
	sqlAddExpectations(mock, test)

	repository := repositories.NewUser()
	tx, err := mockDB.Begin()
	require.NoError(t, err)

	err = repository.Add(tx, test.user)
	if !assert.Equal(t, test.want.err, err) && test.want.err != nil {
		assert.EqualError(t, err, test.want.err.Error())
	}

	err = tx.Commit()
	require.NoError(t, err)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}

func sqlAddExpectations(mock sqlmock.Sqlmock, test *userTestCase) {
	mock.ExpectBegin()

	reqRegExp := &strings.Builder{}
	fmt.Fprint(reqRegExp, "^INSERT[\t \n]+INTO[\t \n]+http_users[\t \n]+")
	fmt.Fprint(reqRegExp, "\\(id, name, lastname, birthdate\\)[\t \n]+")
	fmt.Fprint(reqRegExp, "VALUES\\(DEFAULT,\\$1,\\$2,\\$3\\)[\t \n]*$")

	ex := mock.ExpectExec(reqRegExp.String()).
		WithArgs(test.user.Name, test.user.Lastname, test.user.Birthdate)

	if test.name == "insert fail" {
		ex.WillReturnError(test.txRet.err)
	} else {
		ex.WillReturnResult(test.txRet.result)
	}

	mock.ExpectCommit()
}
