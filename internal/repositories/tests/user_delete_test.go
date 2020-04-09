package repositories_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/goofinator/usersHttp/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var userDeleteTests = []*userTestCase{
	&userTestCase{
		name: "insert fail",
		txRet: userResult{err: errSome,
			result: sqlmock.NewResult(0, 1)},
		want: userResult{err: errSome},
	},
	&userTestCase{
		name: "0 rows",
		txRet: userResult{
			err:    nil,
			result: sqlmock.NewResult(0, 0),
		},
		want: userResult{err: err0LineResult},
	},
	&userTestCase{
		name: "2 rows",
		txRet: userResult{
			err:    nil,
			result: sqlmock.NewResult(0, 2),
		},
		want: userResult{err: err2LineResult},
	},
	&userTestCase{
		name: "success",
		txRet: userResult{
			err:    nil,
			result: sqlmock.NewResult(0, 1),
		},
		want: userResult{err: nil},
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
	sqlDeleteExpectations(mock, test)

	repository := repositories.NewUser()
	tx, err := mockDB.Begin()
	require.NoError(t, err)

	err = repository.Delete(tx, test.id)
	if !assert.Equal(t, test.want.err, err) && test.want.err != nil {
		assert.EqualError(t, err, test.want.err.Error())
	}

	err = tx.Commit()
	require.NoError(t, err)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}

func sqlDeleteExpectations(mock sqlmock.Sqlmock, test *userTestCase) {
	mock.ExpectBegin()

	ex := mock.ExpectExec("DELETE FROM http_users WHERE id=\\$1").
		WithArgs(test.id)

	if test.name == "insert fail" {
		ex.WillReturnError(test.txRet.err)
	} else {
		ex.WillReturnResult(test.txRet.result)
	}

	mock.ExpectCommit()
}
